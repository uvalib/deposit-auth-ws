package edu.virginia.depositauthws.core;

import io.dropwizard.Configuration;
import io.dropwizard.db.DataSourceFactory;

import com.fasterxml.jackson.annotation.JsonProperty;
import org.hibernate.validator.constraints.NotEmpty;

import javax.validation.Valid;
import javax.validation.constraints.NotNull;

public class ServiceConfiguration extends Configuration {

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
            //System.out.println( "*** OVERRIDE DATABASE URL ***" );
            database.setUrl( System.getenv( envDbUrl ) );
        }

        if( System.getenv( envDbUser ) != null ) {
            //System.out.println( "*** OVERRIDE DATABASE USER ***" );
            database.setUser( System.getenv( envDbUser ) );
        }

        if( System.getenv( envDbPassword ) != null ) {
            //System.out.println( "*** OVERRIDE DATABASE PASSWORD ***" );
            database.setPassword( System.getenv( envDbPassword ) );
        }
    }
}
