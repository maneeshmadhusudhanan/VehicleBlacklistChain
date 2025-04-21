



<div align="center">
  <img 
    src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&weight=900&size=28&duration=2000&pause=800&color=FF007F&center=true&vCenter=true&width=600&height=60&lines=🚨+VEHICLE+BLACKLIST+CHAIN+🚨;Secure+and+Tamper-proof+Ledger;Powered+by+Blockchain+Technology" 
    alt="Typing SVG" 
  />
</div>
---


🚗 Project Overview

Title: VehicleBlacklistChain

Tagline: "Tamper-proof vehicle blacklisting and recovery system using Hyperledger Fabric."

✨ Why This Project?

VehicleBlacklistChain provides a decentralized way to report, track, and remove blacklisted vehicles using a permissioned blockchain. It enhances:

✅ Security: Tamper-proof audit history.

✅ Transparency: Clear role-based access for stakeholders.

✅ Automation: Instant event triggers and seamless flow.

📊 Key Highlights

🔗 Blockchain-powered trust with Hyperledger Fabric.

👤 Role-based Access Control (ABAC) using MSP IDs.

✨ Live Event Tracking:

VehicleBlacklistedEvent

VehicleRecoveredEvent

📃 Full Vehicle Audit History using Fabric ledger state and history APIs.

👥 Target Users

Auto Manufacturers

Dealerships

Government Transport Bodies (MVD)

Blockchain Developers / Enthusiasts

🛠 Architecture Overview

graph TD
  A[Manufacturer Org] -->|Registers Vehicles| B[Chaincode]
  C[Dealer Org] -->|Reports Stolen| B
  D[MVD Org] -->|Removes from Blacklist| B
  B --> E[CouchDB Ledger]
  B --> F[Emit Events]

🛠 Prerequisites

Tools Required:

Docker + Docker Compose

Hyperledger Fabric v2.4.8

MiniFab

Go v1.18 or higher

Knowledge Base:

Blockchain & Smart Contracts

Hyperledger Fabric Ecosystem

Basic GoLang

⚙️ Setup & Installation

Launch Fabric Network

minifab netup -s couchdb -e true -i 2.4.8 -o manufacturer.auto.com

Create & Join Channel

minifab create -c autochannel
minifab join -c autochannel
minifab anchorupdate

🧩 Chaincode Deployment

Place VehicleBlacklistContract.go in the /chaincode/ folder.

Deploy using MiniFab or Fabric CLI:

./start_network.sh       # Start Network
./deploy_chaincode.sh    # Deploy Chaincode
node app.js              # Client App

🎬 Demo Walkthrough

Example Flow:

Create Vehicle

peer chaincode invoke -n vehiclecc -c '{"function":"CreateVehicle", "Args":["V123", "ModelX", "Black", "Maneesh", "2025-04-01", "owner-auto-com"]}'

Report as Stolen

peer chaincode invoke -n vehiclecc -c '{"function":"ReportStolen", "Args":["V123", "Reported stolen from parking"]}'

Remove from Blacklist

peer chaincode invoke -n vehiclecc -c '{"function":"RemoveFromBlacklist", "Args":["V123"]}'

✅ Testing & Validation

Write unit tests using mockstub

Validate:

Role-based access with MSPID

Contract logic and ledger updates

Event emissions

📄 License & Credits

Licensed under the MIT License






//////////////////////////////////////////////
