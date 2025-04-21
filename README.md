

<div align="center">
  <img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=900&size=40&duration=3000&pause=1000&color=FF007F&background=FFB3D899&center=true&vCenter=true&width=800&height=80&lines=🚗+VEHICLE+BLACKLIST+CHAIN;%F0%9F%94%90_Trusted+%7C+Immutable+%7C+Decentralized_%F0%9F%94%92" alt="Header">
</div>

<div align="center">
  <br>
  <img src="https://img.shields.io/badge/Blockchain-3DDC84?style=for-the-badge&logo=blockchain-dot-com&logoColor=white">
  <img src="https://img.shields.io/badge/Hyperledger-2F3134?style=for-the-badge&logo=hyperledger&logoColor=white">
  <img src="https://img.shields.io/badge/Smart_Contracts-FF6C37?style=for-the-badge&logo=ethereum&logoColor=white">
</div>




<p align="center">
  <img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=900&size=28&duration=3000&pause=1000&color=FF007F&background=00000000&center=true&vCenter=true&width=700&height=80&lines=🚗+VEHICLE+BLACKLIST+CHAIN" alt="Typing Animation" />
</p>



<div align="center">
  <br>
  <img src="https://img.shields.io/badge/hyperledger-2F3134?style=for-the-badge&logo=hyperledger&logoColor=white">
  <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white">
  <img src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white">
  <img src="https://img.shields.io/badge/CouchDB-EA4C89?style=for-the-badge&logo=couchdb&logoColor=white">
</div>

<div align="center">
  <h3>
    <a href="#-project-overview" style="color: #FF007F; text-decoration: none; margin: 0 10px;">Overview</a> •
    <a href="#-key-features" style="color: #FF007F; text-decoration: none; margin: 0 10px;">Features</a> •
    <a href="#-architecture" style="color: #FF007F; text-decoration: none; margin: 0 10px;">Architecture</a> •
    <a href="#-getting-started" style="color: #FF007F; text-decoration: none; margin: 0 10px;">Setup</a> •
    <a href="#-demo" style="color: #FF007F; text-decoration: none; margin: 0 10px;">Demo</a> •
    <a href="#-license" style="color: #FF007F; text-decoration: none; margin: 0 10px;">License</a>
  </h3>
</div>

---

## 🌟 Project Overview

<p align="center" style="font-size: 18px; color: #FF69B4;">
  <em>"A blockchain-powered solution for immutable vehicle blacklisting and recovery management"</em>
</p>

<div style="background: linear-gradient(145deg, #1a1a1a, #2a2a2a); padding: 20px; border-radius: 15px; box-shadow: 0 4px 8px rgba(255,0,127,0.2);">
  
  <h3 style="color: #FF007F; border-bottom: 2px dashed #FF007F; padding-bottom: 8px;">🚨 Problem Statement</h3>
  
  Traditional vehicle blacklist systems suffer from:
  <ul style="color: #CCCCCC;">
    <li>❌ Centralized control points</li>
    <li>❌ Vulnerability to data tampering</li>
    <li>❌ Lack of transparent audit trails</li>
  </ul>
  
  <h3 style="color: #00FF7F; border-bottom: 2px dashed #00FF7F; padding-bottom: 8px; margin-top: 25px;">💡 Our Solution</h3>
  
  <ul style="color: #CCCCCC;">
    <li>✅ <span style="color: #FFD700;">Decentralized</span> authority using Hyperledger Fabric</li>
    <li>✅ <span style="color: #00BFFF;">Immutable</span> record-keeping with blockchain</li>
    <li>✅ <span style="color: #FF69B4;">Real-time</span> event tracking and notifications</li>
  </ul>

</div>

---

## ✨ Key Features

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(300px, 1fr)); gap: 20px; margin-top: 30px;">

<div style="background: #1a1a1a; padding: 20px; border-radius: 12px; border-left: 4px solid #FF007F;">
  <h3 style="color: #FF007F;">🔐 Role-Based Access</h3>
  <p style="color: #CCCCCC;">Granular permissions through MSP IDs:</p>
  <ul style="color: #AAAAAA;">
    <li>Manufacturers: Vehicle registration</li>
    <li>Dealers: Stolen reporting</li>
    <li>MVD: Blacklist management</li>
  </ul>
</div>

<div style="background: #1a1a1a; padding: 20px; border-radius: 12px; border-left: 4px solid #00FF7F;">
  <h3 style="color: #00FF7F;">📜 Audit Trail</h3>
  <p style="color: #CCCCCC;">Complete history tracking:</p>
  <ul style="color: #AAAAAA;">
    <li>Full vehicle lifecycle records</li>
    <li>SHA-256 hashed transactions</li>
    <li>Queryable through CouchDB</li>
  </ul>
</div>

<div style="background: #1a1a1a; padding: 20px; border-radius: 12px; border-left: 4px solid #FFD700;">
  <h3 style="color: #FFD700;">⚡ Smart Events</h3>
  <p style="color: #CCCCCC;">Real-time notifications:</p>
  <ul style="color: #AAAAAA;">
    <li>VehicleBlacklistedEvent</li>
    <li>VehicleRecoveredEvent</li>
    <li>Custom event listeners</li>
  </ul>
</div>

</div>

---

## 🏗 Architecture

```mermaid
graph TD
  A[Manufacturer] -->|Register Vehicle| B{Blockchain Network}
  C[Dealership] -->|Report Stolen| B
  D[MVD] -->|Update Status| B
  B -->|Store Data| E[(CouchDB Ledger)]
  B -->|Trigger Events| F[Event Hub]
  F --> G[Notification Systems]
  F --> H[Analytics Dashboard]


🚀 Getting Started

Prerequisites

# System Requirements
- Docker 20.10+
- Go 1.18+
- Node.js 16+
- MiniFab 1.3+


Installation

# Clone Repository
git clone https://github.com/yourusername/VehicleBlacklistChain.git

# Start Network
cd VehicleBlacklistChain
./network/start_network.sh

# Deploy Chaincode
minifab ccup -n vehiclecc -l go -v 1.0 -p ""

# Start Application
cd application
npm install
node app.js



📜 License

MIT License

Copyright (c) 2023 Maneesh madhusudhanan.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

<div align="center" style="margin-top: 40px;"> <img src="https://komarev.com/ghpvc/?username=yourusername&label=Project%20Visitors&color=ff007f&style=flat" alt="Visitor counter"> <br> <p style="color: #888;">Made developer MANEESH Madhusudhanan using Hyperledger Fabric</p> <div style="display: flex; justify-content: center; gap: 15px; margin-top: 20px;"> <img src="https://img.icons8.com/color/48/000000/hyperledger.png" width="40"> <img src="https://img.icons8.com/color/48/000000/golang.png" width="40"> <img src="https://img.icons8.com/color/48/000000/docker.png" width="40"> </div> </div>


