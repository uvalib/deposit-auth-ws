package edu.virginia.depositauthws.core;

import edu.virginia.depositauthws.db.DepositAuthDAO;
import io.dropwizard.Application;
import io.dropwizard.setup.Bootstrap;
import io.dropwizard.setup.Environment;
import io.dropwizard.jdbi.DBIFactory;

import org.skife.jdbi.v2.DBI;

import edu.virginia.depositauthws.resources.ServiceResource;
import edu.virginia.depositauthws.health.FsHealthCheck;

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
        final DBI jdbi = factory.build( environment, configuration.getDataSourceFactory( ), "mysql" );

        final DepositAuthDAO depositAuthDAO = jdbi.onDemand( DepositAuthDAO.class );

        final ServiceResource resource = new ServiceResource(
                depositAuthDAO,
                configuration.getDataDirName()
        );

        // register the health checkers
        environment.healthChecks().register( "filesystem", new FsHealthCheck( configuration.getDataDirName( ) ) );

        // register the main resource
        environment.jersey().register(resource);
    }

}