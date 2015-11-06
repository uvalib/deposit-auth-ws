package edu.virginia.depositauthws.models;

import javax.ws.rs.core.Response.Status;
import com.fasterxml.jackson.annotation.JsonProperty;

public class BasicResponse {

    private Status status;

    public BasicResponse() {
        // Jackson deserialization
    }

    public BasicResponse( Status status ) {
        this.status = status;
    }

    @JsonProperty
    public long getStatus() { return status.getStatusCode( ); }

    @JsonProperty
    public String getMessage() {
        return status.getReasonPhrase( );
    }
}
