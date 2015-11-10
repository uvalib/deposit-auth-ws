package edu.virginia.depositauthws.core;

import edu.virginia.depositauthws.db.DepositAuthDAO;
import io.dropwizard.Application;
import io.dropwizard.setup.Bootstrap;
import io.dropwizard.setup.Environment;
import io.dropwizard.jdbi.DBIFactory;

import org.skife.jdbi.v2.DBI;

import edu.virginia.depositauthws.resources.ServiceResource;
import edu.virginia.depositauthws.health.ServiceHealthCheck;

public class ServiceApplication extends Application<ServiceConfiguration> {

    public static void main( String[] args) throws Exception {
        new ServiceApplication().run(args);
    }

    @Override
    public String getName( ) {
        return "depositauth";
    }

    @Override
    public void initialize(Bootstrap<ServiceConfiguration> bootstrap) {
        // nothing to do yet
    }

    @Override
    public void run( ServiceConfiguration configuration,
                     Environment environment ) {

        final DBIFactory factory = new DBIFactory( );
        final DBI jdbi = factory.build( environment, configuration.getDataSourceFactory( ), "h2" );

        final DepositAuthDAO depositAuthDAO = jdbi.onDemand( DepositAuthDAO.class );

        final ServiceResource resource = new ServiceResource(
                depositAuthDAO,
                configuration.getDataDirName()
        );

        final ServiceHealthCheck healthCheck =
                new ServiceHealthCheck( depositAuthDAO, configuration.getDataDirName( ) );

        environment.healthChecks().register( "filesystem", healthCheck );
        environment.jersey().register(resource);
    }

}