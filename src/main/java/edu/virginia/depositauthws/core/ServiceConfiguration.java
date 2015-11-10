package edu.virginia.depositauthws.core;

import io.dropwizard.Configuration;
import io.dropwizard.db.DataSourceFactory;

import com.fasterxml.jackson.annotation.JsonProperty;
import org.hibernate.validator.constraints.NotEmpty;

import javax.validation.Valid;
import javax.validation.constraints.NotNull;

public class ServiceConfiguration extends Configuration {

    @NotEmpty
    private String dataDirName;

    @Valid
    @NotNull
    @JsonProperty( "database" )
    private DataSourceFactory database = new DataSourceFactory();

    public DataSourceFactory getDataSourceFactory() {
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
}
