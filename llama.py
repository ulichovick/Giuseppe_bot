import sys
from llama_cpp import Llama
import argparse


parser = argparse.ArgumentParser()
parser.add_argument("-w", "--word")
args = parser.parse_args()
def llamav2():
    word = args.word
    llm = Llama(model_path="/app/models/llama-2-7b-chat.ggmlv3.q2_K.bin", n_gpu_layers=20)
    template = "[INST] <<SYS>>You are a helpful, respectful and honest assistant.<</SYS>>"
    prompt = word
    prompt = template + prompt + "[/INST]"
    output = llm.create_completion(prompt, echo=True)
    choices = output["choices"]
    choices = choices[0]
    choices = choices["text"].replace(prompt,'')
    sys.stdout.write(str(choices))

if __name__ == "__main__":
    llamav2()