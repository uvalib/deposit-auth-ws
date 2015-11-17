package edu.virginia.depositauthws.core;

import io.dropwizard.Configuration;
import io.dropwizard.db.DataSourceFactory;

import com.fasterxml.jackson.annotation.JsonProperty;
import org.hibernate.validator.constraints.NotEmpty;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.validation.Valid;
import javax.validation.constraints.NotNull;

public class ServiceConfiguration extends Configuration {

    private final static Logger LOG = LoggerFactory.getLogger( ServiceConfiguration.class );

    private static final String envDbUrl = "DATABASE_URL";
    private static final String envDbUser = "DATABASE_USER";
    private static final String envDbPassword = "DATABASE_PASSWORD";

    @NotEmpty
    private String dataDirName;

    @Valid
    @NotNull
    @JsonProperty( "database" )
    private DataSourceFactory database = new DataSourceFactory( );

    public DataSourceFactory getDataSourceFactory() {
        configOverrides( );
        return database;
    }

    @JsonProperty
    public String getDataDirName() {
        return dataDirName;
    }

    @JsonProperty
    public void setDataDirName( String dirname ) {
        this.dataDirName = dirname;
    }

    private void configOverrides( ) {

        // override the configured settings if they are defined in the environment
        if( System.getenv( envDbUrl ) != null ) {
            LOG.info( "DB url from environment: " + System.getenv( envDbUrl ) );
            database.setUrl( System.getenv( envDbUrl ) );
        }

        if( System.getenv( envDbUser ) != null ) {
            LOG.info( "DB user from environment: " + System.getenv( envDbUser ) );
            database.setUser( System.getenv( envDbUser ) );
        }

        if( System.getenv( envDbPassword ) != null ) {
            LOG.info( "DB password from environment: " + System.getenv( envDbPassword ) );
            database.setPassword( System.getenv( envDbPassword ) );
        }
    }
}
