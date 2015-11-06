package edu.virginia.depositauthws.models;

import com.fasterxml.jackson.annotation.JsonProperty;

public class DepositConstraints {

    private String embargo;

    public DepositConstraints() {
        // Jackson deserialization
    }

    public DepositConstraints( String embargo ) {
        this.embargo = embargo;
    }

    @JsonProperty
    public String getEmbargo() {
        return embargo;
    }
}
