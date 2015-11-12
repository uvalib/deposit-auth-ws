package edu.virginia.depositauthws.models;

import com.fasterxml.jackson.annotation.JsonProperty;

public class AuthDetails {

    private String auth;

    public AuthDetails() {
        // Jackson deserialization
    }

    public AuthDetails( String auth ) {
        this.auth = auth;
    }

    @JsonProperty
    public String getAuth() {
        return auth;
    }
}
