package edu.virginia.depositauthws;

import edu.virginia.depositauthws.core.ServicePolicy;
import edu.virginia.depositauthws.db.DepositAuthDAO;
import edu.virginia.depositauthws.models.AuthListResponse;
import edu.virginia.depositauthws.models.BasicResponse;
import edu.virginia.depositauthws.models.DepositAuth;
import edu.virginia.depositauthws.models.DepositConstraints;
import edu.virginia.depositauthws.resources.ServiceResource;

import edu.virginia.depositauthws.sis.SisHelper;
import org.apache.commons.lang3.RandomStringUtils;
import org.apache.commons.lang3.tuple.Pair;

import javax.ws.rs.core.Response;
import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.StandardCopyOption;
import java.nio.file.Paths;

public class TestHelpers {

    //private static final String exampleSisImportFile = SisHelper.sisInputFile( "XXDATEXX" );

    //
    // helper for examining error codes
    //
    public static Boolean responseContains(BasicResponse response, String error ) {
       return( response.getMessage( ).contains( error ) );
    }

    //
    // get a computing Id that can deposit the specified doctype
    //
    public static String getCanDepositComputingId( ServiceResource resource, String doctype ) {
        DepositAuth[] authData = getAuthList( resource );
        for( DepositAuth da : authData ) {
            Pair<Response.Status, DepositConstraints> resDeposit = ServicePolicy.canDeposit( resource.getDAO( ), da.getCid( ), doctype );
            if( resDeposit.getLeft( ).equals( Response.Status.OK ) ) {
                return( da.getCid( ) );
            }
        }
        return( "" );
    }

    //
    // get a known computing Id
    //
    public static String getGoodComputingId( ServiceResource resource ) {
        DepositAuth[] authData = getAuthList( resource );
        if( authData.length != 0 ) {
           return( authData[ 0 ].getCid( ) );
        }
        return( "" );
    }

    //
    // get a known document Id
    //
    public static String getGoodDocumentId( ServiceResource resource ) {
        DepositAuth[] authData = getAuthList( resource );
        for( DepositAuth da : authData ) {
            if( da.getLid( ).isEmpty( ) == false ) return( da.getLid( ) );
        }
        return( "" );
    }

    //
    // get a good doctype
    //
    public static String getGoodDocType( ) {
        return( "PHDDEFENSE" );
    }

    //
    // get a bad (unknown) doc type
    //
    public static String getBadDocType( ) {
        return( getBadId( ) );
    }

    //
    // get a bad Id
    //
    public static String getBadId( ) {
        return( RandomStringUtils.randomAscii( 10 ) );
    }

    //
    // get a good auth token
    //
    public static String getGoodAuthToken( ) {
        return( "super-secret" );
    }

    //
    // get a bad auth token
    //
    public static String getBadAuthToken( ) {
        return( RandomStringUtils.randomAscii( 32 ) );
    }

    //
    // get a good document Id
    //
    public static String getNewDocumentId( ) {
       return( "libra-oa:" + RandomStringUtils.randomNumeric( 5 ) );
    }

    //
    // get a good import date
    //
    public static String getGoodDate( ) {
        return( "150101" );
    }

    //
    // get a bad import date
    //
    public static String getBadDate( ) {
        return( "bad-date" );
    }

    //
    // get a good record identifier
    //
    public static String getGoodRecordId( ServiceResource resource ) {
        DepositAuth[] authData = getAuthList( resource );
        if( authData.length != 0 ) {
            return( authData[ 0 ].getId( ).toString( ) );
        }
        return( "" );
    }

    //
    // get a good but non-existent record identifier
    //
    public static String getMissingRecordId( ) {
        return( "9999999999" );
    }

    //
    // get a bad record identifier
    //
    public static String getBadRecordId( ) {
        return( getBadId( ) );
    }

    //
    // create an example SIS import file
    //
    public static Integer createSisImportFile( String fs, String date ) {
        String src = SisHelper.sisInputFile( "data", "XXDATEXX" );
        String dst = SisHelper.sisInputFile( fs, date );
        try {
            Files.copy(Paths.get(src), Paths.get(dst), StandardCopyOption.REPLACE_EXISTING);
        } catch( IOException ex ) {
            return( 0 );
        }
        return( 38 );   // we know the example file has 38 records
    }

    //
    // count the records in the specified SIS export file
    //
    public static Integer countSisExportFile( String fs, String date ) {
       String exportFile = SisHelper.sisOutputFile( fs, date );
       return( SisHelper.importFromFile( exportFile ).size( ) );
    }

    //
    // count the records in the specified SIS export file
    //
    public static Integer prepareForExport( DepositAuthDAO dao ) {
        return( dao.prepareForExport( 10 ) );
    }

    //
    // count the number of records that are available for export
    //
    public static Integer countSisExportCandidates( DepositAuthDAO dao ) {
        return( dao.getForExport( ).size( ) );
    }

    public static void removeExportFile( String fs, String date ) {
        String exportFile = SisHelper.sisOutputFile( fs, date );
        try {
            Files.delete(Paths.get(exportFile));
        } catch( IOException ex ) {
            // do nothing...
        }
    }

    //
    // get the complete list of deposit authorizations
    //
    private static DepositAuth[] getAuthList( ServiceResource resource ) {
        Response response = resource.allDepositAuth( );
        if( response.getStatus( ) == Response.Status.OK.getStatusCode( ) ) {
            AuthListResponse allAuth = ( AuthListResponse ) response.getEntity( );
            DepositAuth[] authData = allAuth.getData();
            if( authData != null && authData.length != 0 ) return( authData );
        }
        return( new DepositAuth[ 0 ] );
    }
}