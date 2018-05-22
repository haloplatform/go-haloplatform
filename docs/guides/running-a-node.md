# Running A Node

Here is a guide to help you understand how to start a `Glo` node. This will result in a `Full Node`. For instructions on how to change the mode look at [types of nodes](types-of-nodes.md).

## :one: - Clone Github Repo

Run the following commands in your terminal:

```bash
git clone https://github.com/haloplatform/glo.git

cd glo

git checkout testnet_qa
```

## :two: - Install OR Go to Binary Folder

To install the `glo` client you need to copy it to your operating systems command path.

**Linux** :penguin:
```bash
sudo cp pre_built/Ubuntu/geth /usr/local/bin/
```

**Windows** 

Can't easily install in windows, instead just simply `cd` to the correct folder and call the `.exe` directly.

```bash
cd pre_built\Windows\
```

**MacOS**
```bash
sudo cp pre_built/Mac/geth /usr/local/bin/
```

> Note for MacOS users: If your terminal tells you that `/usr/local/bin` doesn't exist run `sudo mkdir /usr/local/bin`

## :three: - Use It

In a terminal run: 
```bash
geth --datadir node_data --port 50608
```

In annother terminal run:

```bash
geth attach node_data/geth.ipc
```



