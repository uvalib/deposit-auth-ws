package edu.virginia.depositauthws.core;

import io.dropwizard.Configuration;
import com.fasterxml.jackson.annotation.JsonProperty;
import org.hibernate.validator.constraints.NotEmpty;

public class ServiceConfiguration extends Configuration {

    @NotEmpty
    private String dataDirName;

    @JsonProperty
    public String getDataDirName() {
        return dataDirName;
    }

    @JsonProperty
    public void setDataDirName( String dirname ) {
        this.dataDirName = dirname;
    }
}
