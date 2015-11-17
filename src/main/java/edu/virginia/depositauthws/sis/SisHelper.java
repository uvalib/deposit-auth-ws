package edu.virginia.depositauthws.sis;

import edu.virginia.depositauthws.core.ServiceHelper;
import edu.virginia.depositauthws.db.DepositAuthDAO;
import edu.virginia.depositauthws.mapper.DepositAuthMapper;
import edu.virginia.depositauthws.models.DepositAuth;
import org.apache.commons.lang3.tuple.Pair;

import javax.ws.rs.core.Response;
import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.List;
import java.util.Date;
import java.util.stream.Stream;
import java.nio.charset.StandardCharsets;

import org.apache.commons.io.FileUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SisHelper {

    private final static Logger LOG = LoggerFactory.getLogger( SisHelper.class );
    private static final String separator = "|";
    private static final SimpleDateFormat sisDateFormatter = new SimpleDateFormat( "MM/dd/yyyy" );

    //
    // write any pending records to SIS
    //
    public static Pair<Response.Status, Integer> exportToSis(DepositAuthDAO depositAuthDAO, String fs, String date ) {

        Integer count = 0;
        List<DepositAuth> depositList = depositAuthDAO.getForExport( );
        if( !depositList.isEmpty( ) ) {
            LOG.info( "Exporting " + depositList.size( ) + " records to SIS" );
            String filename = fs + File.separator + sisOutputFile(date);
            Pair<Response.Status, Integer> status = exportToFile( depositList, filename );
            count = status.getRight( );

            if( ServiceHelper.isOK( status.getLeft( ) ) ) {
                // mark all the exported records as such...
                LOG.info( "Updating exported records" );
                for (DepositAuth da : depositList) {
                    depositAuthDAO.markExported(da.getId().toString());
                }
            } else {
                return( status );
            }

        } else {
            LOG.info( "No suitable records available for SIS" );
        }

        return( Pair.of( Response.Status.OK, count ) );
    }

    //
    // read any pending records from SIS
    //
    public static Pair<Response.Status, Integer> importFromSis( DepositAuthDAO depositAuthDAO, String fs, String date ) {

        Integer count = 0;
        String filename = fs + File.separator + sisInputFile( date );
        if( new File( filename ).isFile( ) ) {
            LOG.info( "Importing from \'" + filename + "\'" );
            List<DepositAuth> depositList = importFromFile( filename );
            if( !depositList.isEmpty( ) ) {
                LOG.info( "Importing " + depositList.size( ) + " records from SIS" );
                for (DepositAuth da : depositList) {
                    if( !importRecord( depositAuthDAO, da ) ) {
                        Pair.of( Response.Status.INTERNAL_SERVER_ERROR, count );
                    }
                    count++;
                }
                try {
                    Files.delete(Paths.get(filename));
                } catch( IOException ex ) {
                    Pair.of( Response.Status.INTERNAL_SERVER_ERROR, count );
                }
            }
        } else {
            LOG.info( "No new SIS file available (" + filename + ")" );
        }
        return( Pair.of( Response.Status.OK, count ) );
    }

    //
    // import SIS records from the specified file
    //
    private static List<DepositAuth> importFromFile( String filename ) {

        List<DepositAuth> imports = new ArrayList<>( );
        try {
            Path p = Paths.get( filename );
            Stream<String> lines = Files.lines(p, StandardCharsets.UTF_8);
            String previousLine = "";
            for( String l : (Iterable<String>) lines::iterator ) {

                //
                // if the record length is 170, assume that the line was truncated
                // and concatinate the record with the next line and use that for the full record
                //

                if( l.length( ) == 170 ) {
                    previousLine = l;
                } else {

                    //System.out.println( "Attempting to convert [" + previousLine + l + "]" );
                    DepositAuth da = fromSis( previousLine + l );
                    previousLine = "";
                    if( da != null ) {
                       //System.out.println("OK");
                       imports.add(da);
                    } else {
                       System.out.println("ERROR");
                    }
                }
            }
        } catch( IOException ex ) {
            imports.clear( );
        }
        return( imports );
    }

    //
    // export the supplied auth records to the specified SIS export file
    //
    private static Pair<Response.Status, Integer> exportToFile( List<DepositAuth> depositList, String filename ) {

        LOG.info( "Exporting to \'" + filename + "\'" );
        File f = new File( filename );
        Integer count = 0;
        try {
            for (DepositAuth da : depositList) {
                String record = toSis(da);
                FileUtils.writeStringToFile( f, record + "\n", "UTF-8", true );
                count++;
            }
        } catch( IOException ex ) {
            Pair.of( Response.Status.INTERNAL_SERVER_ERROR, count );
        }
        return( Pair.of( Response.Status.OK, count ) );
    }

    //
    // import the specified SIS record. Determine if it is an update to an existing record or
    // a new record
    //
    private static Boolean importRecord( DepositAuthDAO depositAuthDAO, DepositAuth da ) {
        Boolean status = true;
        List<DepositAuth> authList = depositAuthDAO.findByCid( da.getCid( ) );
        DepositAuth existing = findExisting( authList, da );
        if( existing == null ) {
            // insert the new item
            status = depositAuthDAO.insert( da ) == 1;
        } else {
            // set the Id and update
            da.setId( existing.getId( ) );
            status = depositAuthDAO.update( da ) == 1;
        }
        return( status );
    }

    //
    // determine if a just received from SIS auth record is an update to an existing record
    //
    private static DepositAuth findExisting( List<DepositAuth> authList, DepositAuth newDa ) {
        for( DepositAuth da : authList ) {
           if( da.getProgram( ).equals( newDa.getProgram( ) ) &&
               da.getDoctype( ).equals( newDa.getDoctype( ) ) ) {
               return( da );
            }
        }
        return( null );
    }
    //
    // generate a SIS record from a deposit auth record
    //
    private static String toSis( DepositAuth da ) {
        String r =
           "emp id" + separator +
           da.getCid( ) + separator +
           "first name" + separator +
           "middle name" + separator +
           "last name" + separator +
           "career" + separator +
           da.getProgram( ) + separator +
           "plan" + separator +
           da.getLid( ) + separator +
           da.getDoctype( ) + separator +
           "degree" + separator +
           toSisDateFormat( da.getApprovedAt( ) );
        return( r );
    }

    //
    // generate a deposit auth record from a SIS record
    //
    private static DepositAuth fromSis( String record ) {

        String[] separated = record.split( "\\" + separator );
        if( separated.length < 12 ) return( null );
        DepositAuth da =  new DepositAuth( );
        // da.xxx( separated[ 0 ] )    // employee ID
        da.setCid( separated[ 1 ] );
        // da.xxx( separated[ 2 ] )    // first name
        // da.xxx( separated[ 3 ] )    // middle name
        // da.xxx( separated[ 4 ] )    // last name
        // da.xxx( separated[ 5 ] )    // career
        da.setProgram( separated[ 6 ] );
        // da.xxx( separated[ 7 ] )    // plan
        da.setTitle( separated[ 8 ] );
        da.setDoctype( separated[ 9 ] );
        // da.xxx( separated[ 10 ] )   // degree
        da.setApprovedAt( toNativeDateFormat( separated[ 11 ] ) );

        return( da );
    }

    //
    // generate the input from SIS filename
    //
    public static String sisInputFile( String date ) {
       return( "UV_Libra_From_SIS_" + date + ".txt" );
    }

    //
    // generate the output to SIS filename
    //
    public static String sisOutputFile( String date ) {
        return( "UV_LIBRA_IN_" + date + ".txt" );
    }

    //
    // convert a date from our native format (YYYY-MM-DD) to the SIS format (MM/DD/YYYY)
    //
    private static String toSisDateFormat( String date ) {
        try {
            Date d = DepositAuthMapper.dateFormat.parse( date );
            return( sisDateFormatter.format( d ) );
        } catch( ParseException ex ) { }
        System.out.println( "ERROR converting " + date + " to SIS format" );
        return( "" );
    }

    //
    // convert a date from the SIS format (MM/DD/YYYY) to our native format (YYYY-MM-DD)
    //
    private static String toNativeDateFormat( String date ) {
        try {
            Date d = sisDateFormatter.parse( date );
            return( DepositAuthMapper.dateFormat.format( d ) );
        } catch( ParseException ex ) { }
        System.out.println( "ERROR converting " + date + " to native format" );
        return( "" );
    }
}
