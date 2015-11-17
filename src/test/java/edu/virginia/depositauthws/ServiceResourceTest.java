package edu.virginia.depositauthws;

import edu.virginia.depositauthws.core.ServiceHelper;
import edu.virginia.depositauthws.models.*;
import io.dropwizard.jdbi.DBIFactory;
import org.skife.jdbi.v2.DBI;
import io.dropwizard.testing.junit.DropwizardAppRule;

import edu.virginia.depositauthws.resources.ServiceResource;
import edu.virginia.depositauthws.db.DepositAuthDAO;
import edu.virginia.depositauthws.core.ServiceApplication;
import edu.virginia.depositauthws.core.ServiceConfiguration;

import org.junit.BeforeClass;
import org.junit.Test;
import org.junit.ClassRule;

import javax.ws.rs.core.Response;

import static org.assertj.core.api.Assertions.assertThat;

public class ServiceResourceTest {

    private static ServiceResource resource;

    //
    // create a running dropwizard app environment
    //
    @ClassRule
    public static DropwizardAppRule<ServiceConfiguration> rule = new DropwizardAppRule<>( ServiceApplication.class,
            "src/main/resources/service.yaml" );

    //
    // do once because it takes time to establish database connections
    //
    @BeforeClass
    public static void once( ) {
        final DBIFactory factory = new DBIFactory( );
        final DBI jdbi = factory.build( rule.getEnvironment( ), rule.getConfiguration( ).getDataSourceFactory( ), "mysql" );
        final DepositAuthDAO depositAuthDAO = jdbi.onDemand( DepositAuthDAO.class );
        resource = new ServiceResource( depositAuthDAO, rule.getConfiguration( ).getDataDirName( ) );
    }

    @Test
    public void getAllAuthList() {
        //
        // get a list of all deposit authorizations
        //
        AuthListResponse authListResponse = resource.allDepositAuth( );
        assertThat( authListResponse.getStatus( ) ).isEqualTo( Response.Status.OK.getStatusCode( ) );

        DepositAuth [] authData = authListResponse.getData( );
        assertThat( authData ).isNotNull( );
        assertThat( authData ).isNotEmpty( );
    }

    @Test
    public void getAuthByGoodComputingId( ) {
        //
        // get a list of deposit authorizations for the specified computing Id
        //
        String id = TestHelpers.getGoodComputingId( resource );
        assertThat( id ).isNotEmpty( );

        AuthListResponse authListResponse = resource.authByComputingId( id );
        assertThat( authListResponse.getStatus( ) ).isEqualTo( Response.Status.OK.getStatusCode( ) );

        DepositAuth [] authData = authListResponse.getData( );
        assertThat( authData ).isNotNull( );
        assertThat( authData ).isNotEmpty( );
    }

    @Test
    public void getAuthByGoodDocumentId( ) {
        //
        // get a list of deposit authorizations containing the specified document Id
        //
        String id = TestHelpers.getGoodDocumentId( resource );
        assertThat( id ).isNotEmpty( );

        AuthListResponse authListResponse = resource.authByDocumentId( id );
        assertThat( authListResponse.getStatus( ) ).isEqualTo( Response.Status.OK.getStatusCode( ) );

        DepositAuth [] authData = authListResponse.getData( );
        assertThat( authData ).isNotNull( );
        assertThat( authData ).isNotEmpty( );
    }

    @Test
    public void getAuthByBadComputingId( ) {
        //
        // ensure we get a NOT FOUND for deposit auths of a non existent computing Id
        //
        String id = TestHelpers.getBadId( );
        assertThat( id ).isNotEmpty( );

        AuthListResponse authListResponse = resource.authByComputingId( id );
        assertThat( authListResponse.getStatus( ) ).isEqualTo( Response.Status.NOT_FOUND.getStatusCode( ) );
        assertThat( TestHelpers.responseContains( authListResponse, ServiceHelper.badCidError ) );
    }

    @Test
    public void getAuthByBadDocumentId( ) {
        //
        // ensure we get a NOT FOUND for deposit auths of a non existent document Id
        //
        String id = TestHelpers.getBadId( );
        assertThat( id ).isNotEmpty( );

        AuthListResponse authListResponse = resource.authByDocumentId( id );
        assertThat( authListResponse.getStatus( ) ).isEqualTo( Response.Status.NOT_FOUND.getStatusCode( ) );
        assertThat( TestHelpers.responseContains( authListResponse, ServiceHelper.badLidError ) );
    }

    @Test
    public void canDepositGoodComputingId( ) {
        //
        // ensure we get an OK when we request to deposit with a good computing Id
        //
        String doctype = TestHelpers.getGoodDocType( );
        String id = TestHelpers.getCanDepositComputingId( resource, doctype );
        assertThat( id ).isNotEmpty( );

        CanDepositResponse canDepositResponse = resource.canDeposit( id, doctype );
        assertThat( canDepositResponse.getStatus( ) ).isEqualTo( Response.Status.OK.getStatusCode( ) );
    }

    @Test
    public void canDepositBadDoctype( ) {
        //
        // ensure we get an FORBIDDEN when we request to deposit with a bad document type
        //
        String doctype = TestHelpers.getGoodDocType( );
        String id = TestHelpers.getCanDepositComputingId( resource, doctype );
        assertThat( id ).isNotEmpty( );

        // set to a bad doctype
        doctype = TestHelpers.getBadDocType( );
        CanDepositResponse canDepositResponse = resource.canDeposit( id, doctype );
        assertThat( canDepositResponse.getStatus( ) ).isEqualTo( Response.Status.FORBIDDEN.getStatusCode( ) );
        assertThat( TestHelpers.responseContains( canDepositResponse, ServiceHelper.badDocTypeError ) );
    }

    @Test
    public void canDepositBadComputingId( ) {
        //
        // ensure we get a FORBIDDEN when we request to deposit with a non existent computing Id
        //
        String id = TestHelpers.getBadId( );
        String doctype = TestHelpers.getGoodDocType( );
        assertThat( id ).isNotEmpty( );

        CanDepositResponse canDepositResponse = resource.canDeposit( id, doctype );
        assertThat( canDepositResponse.getStatus( ) ).isEqualTo( Response.Status.FORBIDDEN.getStatusCode( ) );
        assertThat( TestHelpers.responseContains( canDepositResponse, ServiceHelper.badCidError ) );
    }

    @Test
    public void doDepositGoodComputingId( ) {
        //
        // ensure we get an OK when we attempt to deposit with a good computing Id
        //
        String doctype = TestHelpers.getGoodDocType( );
        String id = TestHelpers.getCanDepositComputingId( resource, doctype );
        assertThat( id ).isNotEmpty( );

        DepositDetails depositDetails = new DepositDetails( TestHelpers.getGoodAuthToken( ), TestHelpers.getNewDocumentId( ) );
        BasicResponse doDepositResponse = resource.doDeposit( id, doctype, depositDetails );
        assertThat( doDepositResponse.getStatus( ) ).isEqualTo( Response.Status.OK.getStatusCode( ) );
    }

    @Test
    public void doDepositBadAuthToken( ) {
        //
        // ensure we get an UNAUTHORIZED when we attempt to deposit with a bad auth token
        //
        String doctype = TestHelpers.getGoodDocType( );
        String id = TestHelpers.getCanDepositComputingId( resource, doctype );
        assertThat( id ).isNotEmpty( );

        DepositDetails depositDetails = new DepositDetails( TestHelpers.getBadAuthToken( ), TestHelpers.getNewDocumentId( ) );
        BasicResponse doDepositResponse = resource.doDeposit( id, doctype, depositDetails );
        assertThat( doDepositResponse.getStatus( ) ).isEqualTo( Response.Status.UNAUTHORIZED.getStatusCode( ) );
        assertThat( TestHelpers.responseContains( doDepositResponse, ServiceHelper.badAuthTokenError ) );
    }

    @Test
    public void doDeleteGoodRecordId( ) {
        //
        // ensure we get an OK when we attempt to delete with a good Id
        //
        String id = TestHelpers.getGoodRecordId( resource );
        assertThat( id ).isNotEmpty( );

        AuthDetails authDetails = new AuthDetails( TestHelpers.getGoodAuthToken( ) );
        BasicResponse doDeleteResponse = resource.doDelete( id.toString( ), authDetails );
        assertThat( doDeleteResponse.getStatus( ) ).isEqualTo( Response.Status.OK.getStatusCode( ) );
    }

    @Test
    public void doDeleteMissingRecordId( ) {
        //
        // ensure we get an OK when we attempt to delete with a good Id
        //
        String id = TestHelpers.getMissingRecordId( );

        AuthDetails authDetails = new AuthDetails( TestHelpers.getGoodAuthToken( ) );
        BasicResponse doDeleteResponse = resource.doDelete( id, authDetails );
        assertThat( doDeleteResponse.getStatus( ) ).isEqualTo( Response.Status.NOT_FOUND.getStatusCode( ) );
        assertThat( TestHelpers.responseContains( doDeleteResponse, ServiceHelper.badIdError ) );
    }

    @Test
    public void doDeleteBadRecordId( ) {
        //
        // ensure we get an OK when we attempt to delete with a good Id
        //
        String id = TestHelpers.getBadRecordId( );

        AuthDetails authDetails = new AuthDetails( TestHelpers.getGoodAuthToken( ) );
        BasicResponse doDeleteResponse = resource.doDelete( id, authDetails );
        assertThat( doDeleteResponse.getStatus( ) ).isEqualTo( Response.Status.BAD_REQUEST.getStatusCode( ) );
        assertThat( TestHelpers.responseContains( doDeleteResponse, ServiceHelper.badIdError ) );
    }

    @Test
    public void doDeleteBadAuthToken( ) {
        //
        // ensure we get an UNAUTHORIZED when we attempt to delete with a bad auth token
        //
        String id = TestHelpers.getGoodRecordId( resource );

        AuthDetails authDetails = new AuthDetails( TestHelpers.getBadAuthToken( ) );
        BasicResponse doDeleteResponse = resource.doDelete( id, authDetails );
        assertThat( doDeleteResponse.getStatus( ) ).isEqualTo( Response.Status.UNAUTHORIZED.getStatusCode( ) );
    }

    @Test
    public void doImportGoodDate( ) {
        //
        // ensure we get an OK when importing with a good date
        //
        String date = TestHelpers.getGoodDate( );
        Integer count = TestHelpers.createSisImportFile( rule.getConfiguration( ).getDataDirName( ), date );

        AuthDetails authDetails = new AuthDetails( TestHelpers.getGoodAuthToken( ) );
        ImportExportResponse doImportResponse = resource.doImport( date, authDetails );
        assertThat( doImportResponse.getStatus( ) ).isEqualTo( Response.Status.OK.getStatusCode( ) );
        assertThat( doImportResponse.getCount( ) ).isEqualTo( count );
    }

    @Test
    public void doImportBadDate( ) {
        //
        // ensure we get a BAD_REQUEST when importing with a bad date
        //
        String date = TestHelpers.getBadDate( );

        AuthDetails authDetails = new AuthDetails( TestHelpers.getGoodAuthToken( ) );
        ImportExportResponse doImportResponse = resource.doImport( date, authDetails );
        assertThat( doImportResponse.getStatus( ) ).isEqualTo( Response.Status.BAD_REQUEST.getStatusCode( ) );
        assertThat( TestHelpers.responseContains( doImportResponse, ServiceHelper.badDateError ) );
    }

    @Test
    public void doImportBadAuthToken( ) {
        //
        // ensure we get an UNAUTHORIZED when importing with a bad auth token
        //
        String date = TestHelpers.getBadDate( );

        AuthDetails authDetails = new AuthDetails( TestHelpers.getBadAuthToken( ) );
        ImportExportResponse doImportResponse = resource.doImport( date, authDetails );
        assertThat( doImportResponse.getStatus( ) ).isEqualTo( Response.Status.UNAUTHORIZED.getStatusCode( ) );
        assertThat( TestHelpers.responseContains( doImportResponse, ServiceHelper.badAuthTokenError ) );
    }

    @Test
    public void doExportGoodDate( ) {
        //
        // ensure we get an OK when exporting with a good date
        //
        String date = TestHelpers.getGoodDate( );

        AuthDetails authDetails = new AuthDetails( TestHelpers.getGoodAuthToken( ) );
        ImportExportResponse doExportResponse = resource.doExport( date, authDetails );
        assertThat( doExportResponse.getStatus( ) ).isEqualTo( Response.Status.OK.getStatusCode( ) );
    }

    @Test
    public void doExportBadDate( ) {
        //
        // ensure we get a BAD_REQUEST when exporting with a bad date
        //
        String date = TestHelpers.getBadDate( );

        AuthDetails authDetails = new AuthDetails( TestHelpers.getGoodAuthToken( ) );
        ImportExportResponse doExportResponse = resource.doExport( date, authDetails );
        assertThat( doExportResponse.getStatus( ) ).isEqualTo( Response.Status.BAD_REQUEST.getStatusCode( ) );
        assertThat( TestHelpers.responseContains( doExportResponse, ServiceHelper.badDateError ) );
    }

    @Test
    public void doExportBadAuthToken( ) {
        //
        // ensure we get an UNAUTHORIZED when exporting with a bad auth token
        //
        String date = TestHelpers.getBadDate( );

        AuthDetails authDetails = new AuthDetails( TestHelpers.getBadAuthToken( ) );
        ImportExportResponse doExportResponse = resource.doExport( date, authDetails );
        assertThat( doExportResponse.getStatus( ) ).isEqualTo( Response.Status.UNAUTHORIZED.getStatusCode( ) );
        assertThat( TestHelpers.responseContains( doExportResponse, ServiceHelper.badAuthTokenError ) );
    }
}