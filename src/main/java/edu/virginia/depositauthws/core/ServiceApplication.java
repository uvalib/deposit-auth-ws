package edu.virginia.depositauthws.core;

import io.dropwizard.Application;
import io.dropwizard.setup.Bootstrap;
import io.dropwizard.setup.Environment;

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

        final ServiceResource resource = new ServiceResource(
                configuration.getDataDirName()
        );

        final ServiceHealthCheck healthCheck =
                new ServiceHealthCheck( configuration.getDataDirName( ) );

        environment.healthChecks().register( "filesystem", healthCheck );
        environment.jersey().register(resource);
    }

}