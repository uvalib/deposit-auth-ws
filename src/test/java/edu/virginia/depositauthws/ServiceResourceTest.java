package edu.virginia.depositauthws;

import edu.virginia.depositauthws.models.*;
import io.dropwizard.jdbi.DBIFactory;
import org.skife.jdbi.v2.DBI;
import io.dropwizard.testing.junit.DropwizardAppRule;

import edu.virginia.depositauthws.resources.ServiceResource;
import edu.virginia.depositauthws.db.DepositAuthDAO;
import edu.virginia.depositauthws.core.ServiceApplication;
import edu.virginia.depositauthws.core.ServiceConfiguration;

import org.junit.Before;
import org.junit.Test;
import org.junit.ClassRule;

import javax.ws.rs.core.Response;

import static org.assertj.core.api.Assertions.assertThat;

public class ServiceResourceTest {

    private ServiceResource resource;

    //
    // create a running dropwizard app environment
    //
    @ClassRule
    public static DropwizardAppRule<ServiceConfiguration> rule = new DropwizardAppRule<>( ServiceApplication.class,
            "src/main/resources/service.yaml" );

    @Before
    public void setup( ) {
        // Before each test, we re-instantiate our resource
        // It is good practice when dealing with a class that
        // contains mutable data to reset it so tests can be ran independently
        // of each other.

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
        // ensure we get an OK when we attempt to deposit with a good computing Id
        //
        String doctype = TestHelpers.getGoodDocType( );
        String id = TestHelpers.getCanDepositComputingId( resource, doctype );
        assertThat( id ).isNotEmpty( );

        DepositDetails depositDetails = new DepositDetails( TestHelpers.getBadAuthToken( ), TestHelpers.getNewDocumentId( ) );
        BasicResponse doDepositResponse = resource.doDeposit( id, doctype, depositDetails );
        assertThat( doDepositResponse.getStatus( ) ).isEqualTo( Response.Status.UNAUTHORIZED.getStatusCode( ) );
    }

    @Test
    public void doImportGoodDate( ) {
        //
        // ensure importing for a good date is successful
        //
        String date = TestHelpers.getGoodDate( );
    }

    @Test
    public void doImportBadDate( ) {
        //
        // ensure importing for a bad date returns an appropriate error
        //
        String date = TestHelpers.getBadDate( );
    }

    @Test
    public void doImportBadFs( ) {
        //
        // ensure importing for a bad FS returns an appropriate error
        //
        String date = TestHelpers.getGoodDate( );
    }

    @Test
    public void doExportGoodDate( ) {
        //
        // ensure exporting for a good date is successful
        //
        String date = TestHelpers.getGoodDate( );
    }

    @Test
    public void doExportBadDate( ) {
        //
        // ensure exporting for a bad date returns an appropriate error
        //
        String date = TestHelpers.getBadDate( );
    }

    @Test
    public void doExportBadFs( ) {
        //
        // ensure importing for a bad FS returns an appropriate error
        //
        String date = TestHelpers.getGoodDate( );
    }
}