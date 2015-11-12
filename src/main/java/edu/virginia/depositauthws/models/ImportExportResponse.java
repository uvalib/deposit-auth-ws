package edu.virginia.depositauthws.models;

import com.fasterxml.jackson.annotation.JsonProperty;

import javax.ws.rs.core.Response.Status;

public class ImportExportResponse extends BasicResponse {

    private Integer count;

    public ImportExportResponse() {
        // Jackson deserialization
    }

    public ImportExportResponse( Status status, String message ) {
        super( status, message );
        this.count = 0;
    }

    public ImportExportResponse( Status status, Integer count ) {
        super( status );
        this.count = count;
    }

    @JsonProperty
    public Integer getCount() {
        return count;
    }
}
