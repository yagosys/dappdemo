import { ethers } from "./ethers-5.6.esm.min.js"
import { abi, contractAddress } from "./constants.js"

const connectButton = document.getElementById("connectButton")
const withdrawButton = document.getElementById("withdrawButton")
const fundButton = document.getElementById("fundButton")
const balanceButton = document.getElementById("balanceButton")
const latestBlockButton = document.getElementById("latestBlockButton")

connectButton.onclick = connect
withdrawButton.onclick = withdraw
fundButton.onclick = fund
balanceButton.onclick = getBalance
latestBlockButton.onclick = getLatestBlockNew

async function getLatestBlockNew() {
    if (typeof window.ethereum !== "undefined") {
        const provider = new ethers.providers.Web3Provider(window.ethereum);
        try {
            // Fetch the latest block's details
            const latestBlock = await provider.getBlockWithTransactions("latest");
            console.log(latestBlock);

            let transactionsDetails = '';
            for (const tx of latestBlock.transactions) {
                // For each transaction, append its details
                transactionsDetails += `
Transaction Hash: ${tx.hash}
From: ${tx.from}
To: ${tx.to || 'Contract Creation'}
Value: ${ethers.utils.formatEther(tx.value)} ETH
---------------------------
`;
            }

            const blockDetails = `
Block Number: ${latestBlock.number}
Timestamp: ${new Date(latestBlock.timestamp * 1000).toLocaleString()}
Miner: ${latestBlock.miner || 'N/A'}
Transactions: ${latestBlock.transactions.length}
Gas Used: ${latestBlock.gasUsed ? latestBlock.gasUsed.toString() : 'N/A'}
Difficulty: ${latestBlock.difficulty ? latestBlock.difficulty.toString() : 'N/A'}
---------------------------
Transactions Details:
${transactionsDetails}
            `;
            console.log(blockDetails);
            // Optionally display this information in the UI rather than an alert for better readability
             window.alert(blockDetails);
        } catch (error) {
            console.error(`Failed to get latest block details: ${error}`);
        }
    } else {
        latestBlockButton.innerHTML = "Please install MetaMask";
    }
}

async function getLatestBlock() {
    if (typeof window.ethereum !== "undefined") {
        const provider = new ethers.providers.Web3Provider(window.ethereum);
        try {
            // Fetch the latest block's details
            const latestBlock = await provider.getBlock("latest");
            console.log(latestBlock);
           const blockDetails = `
                Block Number: ${latestBlock.number}
                Timestamp: ${new Date(latestBlock.timestamp * 1000).toLocaleString()}
                Miner: ${latestBlock.miner || 'N/A'}
                Transactions: ${latestBlock.transactions ? latestBlock.transactions.length : 'N/A'}
                Gas Used: ${latestBlock.gasUsed ? latestBlock.gasUsed.toString() : 'N/A'}
                Difficulty: ${latestBlock.difficulty ? latestBlock.difficulty.toString() : 'N/A'}
            `;
            window.alert(blockDetails);
            // Displaying some of the block's details as an example
        } catch (error) {
            console.error(`Failed to get latest block details: ${error}`);
        }
    } else {
        latestBlockButton.innerHTML = "Please install MetaMask";
    }
}


async function connect() {
  if (typeof window.ethereum !== "undefined") {
    try {
      await ethereum.request({ method: "eth_requestAccounts" })
    } catch (error) {
      console.log(error)
    }
    connectButton.innerHTML = "Connectted"
    const accounts = await ethereum.request({ method: "eth_accounts" })
    window.alert(accounts)
    console.log(accounts)
  } else {
    connectButton.innerHTML = "Please install MetaMask"
  }
}

async function withdraw() {
  console.log(`Withdrawing...`)
  if (typeof window.ethereum !== "undefined") {
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    await provider.send('eth_requestAccounts', [])
    const signer = provider.getSigner()
    const contract = new ethers.Contract(contractAddress, abi, signer)
    try {
      const transactionResponse = await contract.withdraw()
      await listenForTransactionMine(transactionResponse, provider)
      // await transactionResponse.wait(1)
    } catch (error) {
      console.log(error)
    }
  } else {
    withdrawButton.innerHTML = "Please install MetaMask"
  }
}

async function fund() {
  const ethAmount = document.getElementById("ethAmount").value
  console.log(`Funding with ${ethAmount}...`)
  if (typeof window.ethereum !== "undefined") {
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const contract = new ethers.Contract(contractAddress, abi, signer)
    try {
      const transactionResponse = await contract.fund({
        value: ethers.utils.parseEther(ethAmount),
      })
      await listenForTransactionMine(transactionResponse, provider)
    } catch (error) {
      console.log(error)
    }
  } else {
    fundButton.innerHTML = "Please install MetaMask"
  }
}

async function getBalance() {
	  // Select or create an element to display the balance
  let displayElement = document.getElementById('balanceDisplay');
  if (!displayElement) {
    displayElement = document.createElement('div');
    displayElement.id = 'balanceDisplay';
    document.body.appendChild(displayElement);
  }

  if (typeof window.ethereum !== "undefined") {
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    try {
      const balance = await provider.getBalance(contractAddress)
      console.log(ethers.utils.formatEther(balance));
	    const formattedBalance = ethers.utils.formatEther(balance);
	    displayElement.textContent = `Balance: ${formattedBalance} of SmartContract address ${contractAddress}  ETH`;
      window.alert(ethers.utils.formatEther(balance) + " ETH");
    } catch (error) {
      console.log(error)
	    displayElement.textContent = `Error: ${error.message}`;
    }
  } else {
    balanceButton.innerHTML = "Please install MetaMask";
	  displayElement.textContent = "Please install MetaMask";
  }
}

function listenForTransactionMine(transactionResponse, provider) {
    console.log(`Mining ${transactionResponse.hash}`)
    return new Promise((resolve, reject) => {
        try {
            provider.once(transactionResponse.hash, (transactionReceipt) => {
                console.log(
                    `Completed with ${transactionReceipt.confirmations} confirmations. `
                )
                resolve()
            })
        } catch (error) {
            reject(error)
        }
    })
}

