package edu.virginia.depositauthws.models;

import com.fasterxml.jackson.annotation.JsonProperty;
import javax.validation.constraints.NotNull;

public class DepositAuth {

    @NotNull
    @JsonProperty
    private Long id;

    @NotNull
    @JsonProperty
    private String cid;

    @NotNull
    @JsonProperty
    private String doctype;

    @JsonProperty
    private String lid;

    public DepositAuth() {
        // Jackson deserialization
    }

    public DepositAuth( Long id, String cid, String doctype, String lid ) {
        this.id = id;
        this.cid = cid;
        this.doctype = doctype;
        this.lid = lid;
    }

    public Long getId() {
        return id;
    }

    public String getCid() {
        return cid;
    }

    public String getDoctype() {
        return doctype;
    }

    public String getLid() {
        return lid;
    }

    public DepositAuth setId( Long id ) {
        this.id = id;
        return this;
    }

    public DepositAuth setCid( String cid ) {
        this.cid = cid;
        return this;
    }

    public DepositAuth setDoctype( String doctype ) {
        this.doctype = doctype;
        return this;
    }

    public DepositAuth setLid( String lid ) {
        this.lid = lid;
        return this;
    }
}
