package edu.virginia.depositauthws.models;

import javax.ws.rs.core.Response.Status;
import com.fasterxml.jackson.annotation.JsonProperty;

public class BasicResponse {

    private Integer status;
    private String message;

    public BasicResponse() {
        // Jackson deserialization
    }

    public BasicResponse( Status status ) {
        this.status = status.getStatusCode( );
        this.message = status.getReasonPhrase( );
    }

    public BasicResponse( Status status, String message ) {
        this.status = status.getStatusCode( );
        this.message = status.getReasonPhrase( ) + " (" + message + ")";
    }

    @JsonProperty
    public Integer getStatus() { return this.status; }

    @JsonProperty
    public String getMessage() { return this.message; };
}
