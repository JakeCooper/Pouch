# Setup Lambdas

Install apex

## Make a User With AWS IAM

Make sure you have a secret, a access key, and a Region

Make a new user in the [IAM console](https://console.aws.amazon.com/iam/home?region=us-west-2#/home)

Give this user Administrator Access

*Make sure you make it in the us-west-2 region*

## Get the Environment Variables

Use this user to create a access key and a secret access key

*this is from amazon*
```
The only time that you can view or download the secret access keys is when you create the keys. You cannot recover them later. However, you can create new access keys at any time. You must also have permissions to perform the required IAM actions. For more information, see Granting IAM User Permission to Manage Password Policy and Credentials in the IAM User Guide.

Open the IAM console.

In the navigation pane, choose Users.

- Choose your IAM user name (not the check box).

- Choose the Security credentials tab and then choose Create access key.

- To see the new access key, choose Show. Your credentials will look something like this:

    Access key ID: AKIAIOSFODNN7EXAMPLE
    Secret access key: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

## Add the Environment Variables

Now add the keys on the command line.

```bash
export AWS_ACCESS_KEY_ID="replace with your"
export AWS_SECRET_ACCESS_KEY="replace with your secret access key"
export AWS_REGION="us-west-2"
```

## Deploy

Now from the lambdas folder run apex deploy 

in the daemon directory (must be done first)

`go build daemon.go && ./daemon`

In the lambda directory

`apex deploy -E ~/.pouch/settings.json`

