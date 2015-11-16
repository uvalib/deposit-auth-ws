package edu.virginia.depositauthws.health;

import com.codahale.metrics.health.HealthCheck;
import org.apache.commons.lang3.RandomStringUtils;

import java.io.File;

public class FsHealthCheck extends HealthCheck {

    private String dirname;

    public FsHealthCheck( String dirname ) {
        this.dirname = dirname;
    }

    @Override
    protected Result check() throws Exception {

        if( ! new File( dirname ).isDirectory( ) ) return Result.unhealthy( dirname + " not available" );
        String filename = RandomStringUtils.randomNumeric( 10 );
        File f = new File( dirname + File.separator + filename );
        f.deleteOnExit( );
        if( !f.createNewFile( ) ) return Result.unhealthy( dirname + " not writable" );
        return Result.healthy( );
    }
}