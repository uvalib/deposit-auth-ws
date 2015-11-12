package edu.virginia.depositauthws;

import edu.virginia.depositauthws.core.ServicePolicy;
import edu.virginia.depositauthws.db.DepositAuthDAO;
import edu.virginia.depositauthws.models.AuthListResponse;
import edu.virginia.depositauthws.models.DepositAuth;
import edu.virginia.depositauthws.models.DepositConstraints;
import edu.virginia.depositauthws.resources.ServiceResource;
import org.apache.commons.lang3.RandomStringUtils;
import org.apache.commons.lang3.tuple.Pair;

import javax.ws.rs.core.Response;

public class TestHelpers {

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
        return( "2015-01-01" );
    }

    //
    // get a bad import date
    //
    public static String getBadDate( ) {
        return( "xxx" );
    }

    //
    // get the complete list of deposit authorizations
    //
    private static DepositAuth[] getAuthList( ServiceResource resource ) {
        AuthListResponse allAuth = resource.allDepositAuth( );
        if( allAuth.getStatus( ) == Response.Status.OK.getStatusCode( ) ) {
            DepositAuth[] authData = allAuth.getData();
            if( authData != null && authData.length != 0 ) return( authData );
        }
        return( new DepositAuth[ 0 ] );
    }
}