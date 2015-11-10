package edu.virginia.depositauthws.models;

import com.fasterxml.jackson.annotation.JsonProperty;

public class DepositDetails {

    private String auth;
    private String docId;

    public DepositDetails() {
        // Jackson deserialization
    }

    public DepositDetails( String auth, String docId ) {
        this.auth = auth;
        this.docId = docId;
    }

    @JsonProperty
    public String getAuth() {
        return auth;
    }

    @JsonProperty
    public String getDocId() {
        return docId;
    }
}
