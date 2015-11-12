package edu.virginia.depositauthws.models;

import javax.ws.rs.core.Response.Status;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CanDepositResponse extends BasicResponse {

    private DepositConstraints constraints;

    public CanDepositResponse() {
        // Jackson deserialization
    }

    public CanDepositResponse( Status status ) {
        super( status );
        this.constraints = null;
    }

    public CanDepositResponse( Status status, String message ) {
        super( status, message );
        this.constraints = null;
    }

    public CanDepositResponse( Status status, DepositConstraints constraints ) {
        super( status );
        this.constraints = constraints;
    }

    @JsonProperty
    public DepositConstraints getData() {
        return constraints;
    }
}
