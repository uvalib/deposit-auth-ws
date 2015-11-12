package edu.virginia.depositauthws.models;

import com.fasterxml.jackson.annotation.JsonProperty;

public class DepositDetails extends AuthDetails {

    private String docId;

    public DepositDetails() {
        // Jackson deserialization
    }

    public DepositDetails( String auth, String docId ) {
        super( auth );
        this.docId = docId;
    }

    @JsonProperty
    public String getDocId() {
        return docId;
    }
}
