from huggingface_hub import login
from huggingface_hub import HfApi
from huggingface_hub import DatasetFilter, DatasetSearchArguments
from huggingface_hub.repocard import RepoCard


login("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")
card = RepoCard.load("super_glue", "dataset")
print(card.data.to_dict())