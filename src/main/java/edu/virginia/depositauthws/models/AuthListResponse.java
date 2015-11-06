package edu.virginia.depositauthws.models;

import com.fasterxml.jackson.annotation.JsonProperty;

import javax.ws.rs.core.Response.Status;

public class AuthListResponse extends BasicResponse {

    private DepositAuth [] auths;

    public AuthListResponse() {
        // Jackson deserialization
    }

    public AuthListResponse(Status status, DepositAuth [] auths ) {
        super( status );
        this.auths = auths;
    }

    @JsonProperty
    public DepositAuth [] getData() {
        return auths;
    }
}
