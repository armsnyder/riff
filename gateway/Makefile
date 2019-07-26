.PHONY: deploy
deploy:
	gcloud functions deploy gateway --runtime go111 --region us-central1 --trigger-http --entry-point Handler --project riff-247802 --service-account gateway@riff-247802.iam.gserviceaccount.com
