package edu.virginia.depositauthws.models;

import com.fasterxml.jackson.annotation.JsonProperty;

public class DepositAuth {

    private String cid;

    public DepositAuth() {
        // Jackson deserialization
    }

    public DepositAuth( String cid ) {
        this.cid = cid;
    }

    @JsonProperty
    public String getCid() {
        return cid;
    }
}
