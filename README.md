# VEHICLE BLACKLIST CHAIN
Project: Hyperledger Fabric Network


<div align="center"><img src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&duration=3000&pause=1000&color=00F000&width=435&lines=VehicleBlacklistChain" alt="Typing SVG" />

  

</div>
---

🚗 Project Overview

Title: VehicleBlacklistChain
Tagline: "Tamper-proof vehicle blacklisting and recovery system using Hyperledger Fabric."

Why this Project?

> Ensure secure, transparent, and decentralized blacklisting of stolen vehicles. The platform facilitates seamless interaction between Manufacturers, Dealers, and Motor Vehicle Departments (MVD), leveraging blockchain's immutability.



Key Highlights

Blockchain-powered trust with Hyperledger Fabric.

Role-based access using MSP IDs.

Live event tracking with VehicleBlacklistedEvent and VehicleRecoveredEvent.

Full vehicle audit history using Fabric’s ledger.


Who is this for?

Auto Manufacturers

Dealerships

Government Transport Bodies (MVD)

Blockchain Developers / Enthusiasts



---

🧱 Architecture Overview

graph TD
  A[Manufacturer Org] -->|Registers Vehicles| B[Chaincode]
  C[Dealer Org] -->|Reports Stolen| B
  D[MVD Org] -->|Removes from Blacklist| B
  B --> E[CouchDB Ledger]
  B --> F[Emit Events]




---

🛠️ Prerequisites

Tools:

Docker + Docker Compose

Hyperledger Fabric v2.4.8

MiniFab

Go (v1.18 or above)


Knowledge:

Blockchain & Smart Contracts

Hyperledger Fabric Ecosystem

GoLang Basics



---

⚙️ Setup & Installation

# Launch Fabric Network
minifab netup -s couchdb -e true -i 2.4.8 -o manufacturer.auto.com

# Create & Join Channel
minifab create -c autochannel
minifab join -c autochannel
minifab anchorupdate

🧩 Chaincode Deployment

Place your VehicleBlacklistContract inside chaincode/.

Deploy using MiniFab or Fabric CLI commands.



---

⚡ Quick Start

./start_network.sh         # Start Hyperledger Fabric Network
./deploy_chaincode.sh     # Deploy Chaincode
node app.js               # Run your client app


---

🎬 Demo Walkthrough

Example Flow:

1. Create Vehicle



peer chaincode invoke -n vehiclecc -c '{"function":"CreateVehicle", ...}'

2. Report Stolen



peer chaincode invoke -n vehiclecc -c '{"function":"ReportStolen", ...}'

3. Recover Vehicle



peer chaincode invoke -n vehiclecc -c '{"function":"RemoveFromBlacklist", ...}'


---

✅ Testing & Validation

Write unit tests using mockstub

Validate:

Access control via MSP

Chaincode functions

Event emission and ledger updates





© License & Credits

Licensed under the MIT License.
Special thanks to:

Hyperledger Fabric Core Team

MiniFab Developers

Open Source Contributors

