package edu.virginia.depositauthws.resources;

import edu.virginia.depositauthws.core.ServiceHelper;
import edu.virginia.depositauthws.core.ServicePolicy;
import edu.virginia.depositauthws.models.BasicResponse;
import edu.virginia.depositauthws.models.CanDepositResponse;
import edu.virginia.depositauthws.models.AuthListResponse;
import edu.virginia.depositauthws.models.ImportExportResponse;

import edu.virginia.depositauthws.models.DepositAuth;
import edu.virginia.depositauthws.models.DepositConstraints;
import edu.virginia.depositauthws.models.DepositDetails;
import edu.virginia.depositauthws.models.AuthDetails;

import edu.virginia.depositauthws.db.DepositAuthDAO;

import com.codahale.metrics.annotation.Timed;

import javax.ws.rs.GET;
import javax.ws.rs.POST;
import javax.ws.rs.DELETE;

import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.Consumes;
import javax.ws.rs.PathParam;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

import edu.virginia.depositauthws.sis.SisHelper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;
import org.apache.commons.lang3.tuple.Pair;

@Path( "/depositauth" )
@Produces( MediaType.APPLICATION_JSON )
public class ServiceResource {

    private final static Logger LOG = LoggerFactory.getLogger( ServiceResource.class );

    private final String dirname;
    private final DepositAuthDAO depositAuthDAO;

    public ServiceResource( DepositAuthDAO depositAuthDAO, String dirname ) {
        this.depositAuthDAO = depositAuthDAO;
        this.dirname = dirname;
    }

    //
    // helpers for testing...
    //
    public String getDirName( ) {
        return( this.dirname );
    }

    public DepositAuthDAO getDAO( ) {
        return( this.depositAuthDAO );
    }

    @GET
    @Path( "/" )
    @Timed
    //
    // Get all known deposit authorizations
    //
    public AuthListResponse allDepositAuth( ) {
        List<DepositAuth> depositAuth = depositAuthDAO.getAll( );
        return new AuthListResponse( depositAuth.isEmpty( ) ? Response.Status.NOT_FOUND : Response.Status.OK,
                depositAuth.toArray( new DepositAuth[ 0 ] ) );
    }

    @GET
    @Path( "/cid/{cid}/candeposit/{doctype}" )
    @Timed
    //
    // Can the specified computing Id deposit a document of the specified type
    //
    public CanDepositResponse canDeposit( @PathParam( "cid" ) String cid, @PathParam( "doctype" ) String doctype ) {

        // validate inbound parameters
        Pair<Response.Status, String> validate = ServiceHelper.validateCanDepositRequest( cid, doctype );
        if( !ServiceHelper.isValid( validate.getLeft( ) ) ) {
            return new CanDepositResponse( validate.getLeft( ), validate.getRight( ) );
        }

        LOG.info( "Checking deposit authorization; cid: " + cid + ", doctype: " + doctype );

        // check to see if we can deposit
        Pair<Response.Status, DepositConstraints> res = ServicePolicy.canDeposit( depositAuthDAO, cid, doctype );
        return new CanDepositResponse( res.getLeft( ), res.getRight( ) );
    }

    @GET
    @Path( "/cid/{cid}" )
    @Timed
    //
    // Get the deposit authorizations for the specified computing Id
    //
    public AuthListResponse authByComputingId( @PathParam( "cid" ) String cid ) {
        List<DepositAuth> depositAuth = depositAuthDAO.findByCid( cid );
        return new AuthListResponse( depositAuth.isEmpty( ) ? Response.Status.NOT_FOUND : Response.Status.OK,
                depositAuth.toArray( new DepositAuth[ 0 ] ) );
    }

    @GET
    @Path( "/lid/{lid}" )
    @Timed
    //
    // Get the deposit authorizations for the specified computing Id
    //
    public AuthListResponse authByDocumentId( @PathParam( "lid" ) String lid ) {
        List<DepositAuth> depositAuth = depositAuthDAO.findByLid( lid );
        return new AuthListResponse( depositAuth.isEmpty( ) ? Response.Status.NOT_FOUND : Response.Status.OK,
                depositAuth.toArray( new DepositAuth[ 0 ] ) );
    }

    @POST
    @Path( "/cid/{cid}/deposit/{doctype}" )
    @Consumes( MediaType.APPLICATION_JSON )
    @Timed
    //
    // Do the deposit for the specified computing Id
    //
    public BasicResponse doDeposit( @PathParam( "cid" ) String cid, @PathParam( "doctype" ) String doctype, DepositDetails details ) {

        // validate inbound parameters
        Pair<Response.Status, String> validate = ServiceHelper.validateDoDepositRequest( cid, doctype, details );
        if( !ServiceHelper.isValid( validate.getLeft( ) ) ) {
            return new BasicResponse( validate.getLeft( ), validate.getRight( ) );
        }

        // check that they can deposit
        Pair<Response.Status, DepositConstraints> can = ServicePolicy.canDeposit( depositAuthDAO, cid, doctype );
        if( !ServiceHelper.isValid( can.getLeft( ) ) ) {
            return new CanDepositResponse( can.getLeft( ) );
        }

        return new BasicResponse( Response.Status.OK );
    }

    @DELETE
    @Path( "/id/{id}" )
    @Consumes( MediaType.APPLICATION_JSON )
    @Timed
    //
    // Delete the specified record Id
    //
    public BasicResponse doDelete( @PathParam( "id" ) String id, AuthDetails details ) {

        // validate inbound parameters
        Pair<Response.Status, String> validate = ServiceHelper.validateDoDeleteRequest( id, details );
        if( !ServiceHelper.isValid( validate.getLeft( ) ) ) {
            return new BasicResponse( validate.getLeft( ), validate.getRight( ) );
        }

        // do the delete...
        int status = depositAuthDAO.delete( id );
        return new BasicResponse( status == 1 ? Response.Status.OK : Response.Status.NOT_FOUND );
    }

    @POST
    @Path( "/import/{date}" )
    @Consumes( MediaType.APPLICATION_JSON )
    @Timed
    //
    // Import any new SIS records for the specified date
    //
    public ImportExportResponse doImport( @PathParam( "date" ) String date, AuthDetails details ) {

        // validate inbound parameters
        Pair<Response.Status, String> validate = ServiceHelper.validateImportRequest( date, details );
        if( !ServiceHelper.isValid( validate.getLeft( ) ) ) {
            return new ImportExportResponse( validate.getLeft( ), validate.getRight( ) );
        }

        // do the import...
        Pair<Response.Status, Integer> res = SisHelper.importFromSis( depositAuthDAO, dirname, date );
        LOG.info( "Import status: " + res.getLeft( ) + ", count: " + res.getRight( ) );
        return new ImportExportResponse( res.getLeft( ), res.getRight( ) );
    }

    @POST
    @Path( "/export/{date}" )
    @Consumes( MediaType.APPLICATION_JSON )
    @Timed
    //
    // Import any new SIS records for the specified date
    //
    public ImportExportResponse doExport( @PathParam( "date" ) String date, AuthDetails details ) {

        // validate inbound parameters
        Pair<Response.Status, String> validate = ServiceHelper.validateExportRequest( date, details );
        if( !ServiceHelper.isValid( validate.getLeft( ) ) ) {
            return new ImportExportResponse( validate.getLeft( ), validate.getRight( ) );
        }

        // do the export...
        Pair<Response.Status, Integer> res = SisHelper.exportToSis( depositAuthDAO, dirname, date );
        LOG.info( "Export status: " + res.getLeft( ) + ", count: " + res.getRight( ) );
        return new ImportExportResponse( res.getLeft( ), res.getRight( ) );
    }
}