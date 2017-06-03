// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";

import {VPC} from "./vpc";

export class SecurityGroup extends lumi.Resource implements SecurityGroupArgs {
    public readonly name: string;
    public readonly groupDescription: string;
    public readonly groupName?: string;
    public readonly vpc?: VPC;
    public securityGroupEgress?: SecurityGroupRule[];
    public securityGroupIngress?: SecurityGroupRule[];
    @lumi.out public groupID: string;

    constructor(name: string, args: SecurityGroupArgs) {
        super();
        if (name === undefined) {
            throw new Error("Missing required resource name");
        }
        this.name = name;
        if (args.groupDescription === undefined) {
            throw new Error("Missing required argument 'groupDescription'");
        }
        this.groupDescription = args.groupDescription;
        this.groupName = args.groupName;
        this.vpc = args.vpc;
        this.securityGroupEgress = args.securityGroupEgress;
        this.securityGroupIngress = args.securityGroupIngress;
    }
}

export interface SecurityGroupArgs {
    readonly groupDescription: string;
    readonly groupName?: string;
    readonly vpc?: VPC;
    securityGroupEgress?: SecurityGroupRule[];
    securityGroupIngress?: SecurityGroupRule[];
}

export interface SecurityGroupRule {
    ipProtocol: string;
    cidrIp?: string;
    fromPort?: number;
    toPort?: number;
}


