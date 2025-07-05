Functional Technical Document – Shadownet Project

1. Project Name

Shadownet – A distributed, encrypted system for file synchronization and transfer without centralized servers.

2. General Objective

To develop a private, modular network of trusted nodes that enables secure file sharing, synchronization, and protection between devices and users without relying on third-party servers or cloud storage.

3. Core Principles

Distributed network of nodes identified by public keys

Mandatory end-to-end encryption

No centralized server or global database

No user accounts: each device (node) is its own identity

Direct transfers between nodes or encrypted deposit if offline

Local password protection for sensitive directories or nodes

Dual access: powerful CLI and friendly local GUI (localhost)

Multi-drive support: run multiple logical environments on the same physical machine

4. Use Cases

Synchronize files across personal devices

Share files securely between trusted users

Use a shared NAS between multiple users without exposing data

Leave files for delivery to offline devices

Protect specific directories within a shared node

View and manage trusted nodes from a local interface

5. General Architecture

Each node has its own RSA key pair

Local encrypted storage directory (drive/)

Nodes connect directly using secure protocols

All actions are signed and verifiable

Relay functionality if a node trusts both sender and receiver

Physical devices can host multiple logical “drives,” each with its own configuration, keys, and trusted nodes

6. Project Structure

shadownet/
├── cmd/              # CLI interface (Go + cobra)
├── core/             # Core logic: network, crypto, storage
├── internal/         # Daemons, pairing, folder protection
├── drives/           # Drive management and multi-instance logic
├── web/              # GUI (Svelte or React)
├── docs/             # Technical and functional documentation
├── main.go
├── go.mod
├── Makefile
├── build.sh

7. Action Plan (Technical Roadmap)

Phase 0 – Project Base



Phase 1 – Node Identity and Base Config



Phase 2 – Pairing and Trusted Node Management



Phase 3 – File Transfer



Phase 4 – CLI and GUI



Phase 5 – Local Security and Folder Protection



Phase 6 – Multi-drive Architecture



Phase 7 – Packaging and Release



8. Risks and Considerations

Sharing a physical node across multiple drives must enforce strict isolation
→ Mitigation: fully separate folders, processes, and config scopes

Public key injection
→ Mitigation: manual fingerprint verification

Private key loss
→ Mitigation: encrypted backup + optional password protection

9. License and Philosophy

Open-source with a privacy-first mindset

Designed for users who want to escape centralized ecosystems

Modular, extensible, and adaptable to custom environments

This document will serve as a living reference and will be updated as the project evolves.

