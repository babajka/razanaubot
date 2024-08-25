# Razanau Bot

#### Setup local development

* Create `local.yaml` (for local runs) and `production.yaml` files in `secrets/`
with the same structure as `local_base.yaml`. Look up necessary credentials in secure storages.

#### Server setup notes

* You need Redis to be running on the server.

#### Server cronjob

Edit with `crontab -e`

Current config (note that server time is UTC):

```
1 6 * * * /home/wir-prod/deployed/razanaubot
```
