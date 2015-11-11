package edu.virginia.depositauthws;

import edu.virginia.depositauthws.models.AuthListResponse;
import edu.virginia.depositauthws.models.DepositAuth;
import edu.virginia.depositauthws.resources.ServiceResource;
import org.apache.commons.lang3.RandomStringUtils;

import javax.ws.rs.core.Response;

public class TestHelpers {

    public static String getGoodComputingId( ServiceResource resource ) {
        DepositAuth[] authData = getAuthList( resource );
        if( authData.length != 0 ) {
           return( authData[ 0 ].getCid( ) );
        }
        return( "" );
    }

    public static String getGoodDocumentId( ServiceResource resource ) {
        DepositAuth[] authData = getAuthList( resource );
        for( DepositAuth da : authData ) {
            if( da.getLid( ).isEmpty( ) == false ) return( da.getLid( ) );
        }
        return( "" );
    }

    public static String getGoodDocType( ) {
        return( RandomStringUtils.random( 10 ) );
    }

    public static String getBadId( ) {
       return( RandomStringUtils.random( 10 ) );
    }

    private static DepositAuth[] getAuthList( ServiceResource resource ) {
        AuthListResponse allAuth = resource.allDepositAuth( );
        if( allAuth.getStatus( ) == Response.Status.OK.getStatusCode( ) ) {
            DepositAuth[] authData = allAuth.getData();
            if( authData != null && authData.length != 0 ) return( authData );
        }
        return( new DepositAuth[ 0 ] );
    }
}