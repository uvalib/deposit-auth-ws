package edu.virginia.depositauthws.resources;

import edu.virginia.depositauthws.models.BasicResponse;
import edu.virginia.depositauthws.models.CanDepositResponse;
import edu.virginia.depositauthws.models.AuthListResponse;
import edu.virginia.depositauthws.models.DepositAuth;
import edu.virginia.depositauthws.models.DepositConstraints;
import edu.virginia.depositauthws.models.DepositDetails;

import edu.virginia.depositauthws.db.DepositAuthDAO;

import com.codahale.metrics.annotation.Timed;

import javax.ws.rs.GET;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.Consumes;
import javax.ws.rs.PathParam;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path( "/depositauth" )
@Produces( MediaType.APPLICATION_JSON )
public class ServiceResource {

    private final String dirname;
    private final DepositAuthDAO depositAuthDAO;

    public ServiceResource( DepositAuthDAO depositAuthDAO, String dirname ) {
        this.depositAuthDAO = depositAuthDAO;
        this.dirname = dirname;
    }

    @GET
    @Path( "/" )
    @Timed
    //
    // Get all known deposit authorizations
    //
    public AuthListResponse allDepositAuth( ) {
        return new AuthListResponse( Response.Status.OK, depositAuthDAO.getAll( ).toArray( new DepositAuth[ 0 ]) );
    }

    @GET
    @Path( "/cid/{cid}/candeposit/{doctype}" )
    @Timed
    //
    // Can the specified computing Id deposit a document of the specified type
    //
    public CanDepositResponse canDeposit( @PathParam( "cid" ) String cid, @PathParam( "doctype" ) String doctype ) {
        return new CanDepositResponse( Response.Status.OK, new DepositConstraints( ) );
    }

    @GET
    @Path( "/cid/{cid}" )
    @Timed
    //
    // Get the deposit authorizations for the specified computing Id
    //
    public AuthListResponse authByComputingId( @PathParam( "cid" ) String cid ) {
        return new AuthListResponse( Response.Status.OK, depositAuthDAO.findByCid( cid ).toArray( new DepositAuth[ 0 ] ) );
    }

    @GET
    @Path( "/lid/{lid}" )
    @Timed
    //
    // Get the deposit authorizations for the specified computing Id
    //
    public AuthListResponse authByDocumentId( @PathParam( "lid" ) String lid ) {
        return new AuthListResponse( Response.Status.OK, depositAuthDAO.findByLid( lid ).toArray( new DepositAuth[ 0 ] ) );
    }

    @POST
    @Path( "/cid/{cid}/deposit/{doctype}" )
    @Consumes( MediaType.APPLICATION_JSON )
    @Timed
    //
    // Get the deposit authorizations for the specified computing Id
    //
    public BasicResponse didDeposit( @PathParam( "cid" ) String cid, @PathParam( "lid" ) String lid, DepositDetails details ) {
        //DepositAuth = depositAuthDAO.findByCidAndDoctype( cid, xxx );

        return new BasicResponse( Response.Status.OK );
    }
}