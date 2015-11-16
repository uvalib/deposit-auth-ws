package edu.virginia.depositauthws.sis;

import edu.virginia.depositauthws.db.DepositAuthDAO;
import edu.virginia.depositauthws.models.DepositAuth;
import org.apache.commons.lang3.tuple.Pair;

import javax.ws.rs.core.Response;
import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

import java.util.List;
import java.util.stream.Stream;
import java.nio.charset.StandardCharsets;

import org.apache.commons.io.FileUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SisHelper {

    private final static Logger LOG = LoggerFactory.getLogger( SisHelper.class );
    private static final String seperator = "|";

    //
    // write any pending records to SIS
    //
    public static Pair<Response.Status, Integer> exportToSis(DepositAuthDAO depositAuthDAO, String fs, String date ) {

        List<DepositAuth> depositList = depositAuthDAO.getForExport( );
        if( !depositList.isEmpty( ) ) {
            LOG.info( "Exporting " + depositList.size( ) + " records to SIS" );
            String filename = fs + File.separator + outputFile(date);
            return( exportToFile( depositList, depositAuthDAO, filename ) );
        } else {
            LOG.info( "No suitable records available for SIS" );
        }

        return( Pair.of( Response.Status.OK, 0 ) );
    }

    //
    // read any pending records from SIS
    //
    public static Pair<Response.Status, Integer> importFromSis( DepositAuthDAO depositAuthDAO, String fs, String date ) {

        String filename = fs + File.separator + inputFile( date );
        if( new File( filename ).isFile( ) ) {
            LOG.info( "Importing from \'" + filename + "\'" );
            return( importFromFile( depositAuthDAO, filename ) );
        } else {
            LOG.info( "No new SIS file available (" + filename + ")" );
        }
        return( Pair.of( Response.Status.OK, 0 ) );
    }

    private static Pair<Response.Status, Integer> importFromFile( DepositAuthDAO depositAuthDAO, String filename ) {

        Integer count = 0;
        try {
            Path p = Paths.get( filename );
            Stream<String> lines = Files.lines(p, StandardCharsets.UTF_8);
            for( String line : (Iterable<String>) lines::iterator ) {
                if( !importRecord( depositAuthDAO, line ) ) return( Pair.of( Response.Status.INTERNAL_SERVER_ERROR, count ) );
                count++;
            }
            Files.delete( p );
        } catch( IOException ex ) {
            Pair.of( Response.Status.INTERNAL_SERVER_ERROR, 0 );
        }
        return( Pair.of( Response.Status.OK, count ) );
    }

    private static Pair<Response.Status, Integer> exportToFile( List<DepositAuth> depositList, DepositAuthDAO depositAuthDAO, String filename ) {

        LOG.info( "Exporting to \'" + filename + "\'" );
        File f = new File( filename );
        Integer count = 0;
        try {
            for (DepositAuth da : depositList) {
                String record = toSis(da);
                FileUtils.writeStringToFile( f, record + "\n", "UTF-8", true );
                depositAuthDAO.markExported( da.getId( ).toString( ) );
                count++;
            }
        } catch( IOException ex ) {
            Pair.of( Response.Status.INTERNAL_SERVER_ERROR, count );
        }
        return( Pair.of( Response.Status.OK, count ) );
    }

    private static Boolean importRecord( DepositAuthDAO depositAuthDAO, String record ) {
        DepositAuth da = fromSis( record );
        if( da != null ) {
            //depositAuthDAO.
        }
        return( true );
    }

    //
    // generate a SIS record from a deposit auth record
    //
    private static String toSis( DepositAuth da ) {
        String r =
           "emp id" + seperator +
           da.getCid( ) + seperator +
           "first name" + seperator +
           "middle name" + seperator +
           "last name" + seperator +
           "career" + seperator +
           da.getProgram( ) + seperator +
           "plan" + seperator +
           da.getLid( ) + seperator +
           da.getDoctype( ) + seperator +
           "degree" + seperator +
           da.getApprovedAt( );
        return( r );
    }

    //
    // generate a deposit auth record from a SIS record
    //
    private static DepositAuth fromSis( String record ) {

        String[] separated = record.split( "\\" + seperator );
        if( separated.length < 12 ) return( null );
        DepositAuth da =  new DepositAuth( );
        // da.xxx( seperated[ 0 ] )    // employee ID
        da.setCid( separated[ 1 ] );
        // da.xxx( seperated[ 2 ] )    // first name
        // da.xxx( seperated[ 3 ] )    // middle name
        // da.xxx( seperated[ 4 ] )    // last name
        // da.xxx( seperated[ 5 ] )    // career
        da.setProgram( separated[ 6 ] );
        // da.xxx( seperated[ 7 ] )    // plan
        da.setTitle( separated[ 8 ] );
        da.setDoctype( separated[ 9 ] );
        // da.xxx( seperated[ 10 ] )   // degree
        da.setApprovedAt( separated[ 11 ] );

        return( da );
    }

    //
    // generate the input from SIS filename
    //
    private static String inputFile( String date ) {
       return( "UV_Libra_From_SIS" + date + ".txt" );
    }

    //
    // generate the output to SIS filename
    //
    private static String outputFile( String date ) {
        return( "UV_LIBRA_IN" + date + ".txt" );
    }
}
