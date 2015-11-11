package edu.virginia.depositauthws.core;

import edu.virginia.depositauthws.db.DepositAuthDAO;
import edu.virginia.depositauthws.models.DepositAuth;
import edu.virginia.depositauthws.models.DepositConstraints;

import java.util.List;

import edu.virginia.depositauthws.models.DepositDetails;
import org.apache.commons.lang3.tuple.Pair;

import javax.ws.rs.core.Response;

public class ServicePolicy {

    public static Pair<Response.Status, DepositConstraints> canDeposit(DepositAuthDAO depositAuthDAO, String cid, String doctype ) {

        // find any erxisting deposit authorizations
        List<DepositAuth> depositAuth = depositAuthDAO.findByCidAndDoctype( cid, doctype );

        // no deposit information located... send a FORBIDDEN
        if( depositAuth.isEmpty( ) ) {
            return( Pair.of( Response.Status.FORBIDDEN, null ) );
        }

        for( DepositAuth da : depositAuth ) {
            if( da.getDoctype( ).equals( doctype ) ) {
                return( Pair.of( Response.Status.OK, new DepositConstraints( ) ) );
            }
        }
        return( Pair.of( Response.Status.FORBIDDEN, null ) );
    }

    public static Pair<Response.Status, String> checkDeposit( DepositDetails details ) {
        return( Pair.of( Response.Status.FORBIDDEN, null ) );
    }

}
