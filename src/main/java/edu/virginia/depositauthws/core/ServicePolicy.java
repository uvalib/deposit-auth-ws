package edu.virginia.depositauthws.core;

import edu.virginia.depositauthws.db.DepositAuthDAO;
import edu.virginia.depositauthws.models.DepositAuth;
import edu.virginia.depositauthws.models.DepositConstraints;

import java.util.List;

import org.apache.commons.lang3.tuple.Pair;

import javax.ws.rs.core.Response;

public class ServicePolicy {

    public static Pair<Response.Status, DepositConstraints> canDeposit(DepositAuthDAO depositAuthDAO, String cid, String doctype ) {

        // find any existing deposit authorizations
        List<DepositAuth> depositAuth = depositAuthDAO.findByCidAndDoctype( cid, doctype );

        for( DepositAuth da : depositAuth ) {
            if( da.getDoctype( ).equals( doctype ) ) {
                return( Pair.of( Response.Status.OK, new DepositConstraints( ) ) );
            }
        }
        return( Pair.of( Response.Status.FORBIDDEN, null ) );
    }
}
