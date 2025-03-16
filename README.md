## CNAI CLI

Highly simplified way of converting your AI models to AI Artifacts that comply with the [CNAI spec](https://github.com/CloudNativeAI/model-spec/blob/main/docs/spec.md).

## Motivation
- I need a simple way to package AI models as an OCI artifact.

## Goals
- [x] upload any model on huggingface as an AI artifact

## Usage

Insatll

```sh
go install github.com/bupd/cnai-cli/cmd/cnai-cli@latest
```

Convert Models from huggingface to AI Artifact that you can distribute & deploy.

```sh
cnai-cli "hugging_face_git_clone_url" "oci_registry_url"
```
