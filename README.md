# Atlas Mesh

> Service registration and middleware generation SDK for x402 ecosystem

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![x402](https://img.shields.io/badge/x402-Compatible-green)](https://x402.org)

Atlas Mesh enables developers to register x402-protected services and generate middleware code.

## Installation

### TypeScript/JavaScript

```bash
npm install @atlas402/mesh
```

### Python

```bash
pip install atlas-mesh
```

### Go

```bash
go get github.com/atlas402/mesh
```

### Java

```xml
<dependency>
  <groupId>com.atlas402</groupId>
  <artifactId>mesh</artifactId>
  <version>1.0.0</version>
</dependency>
```

## Quick Start

### TypeScript

```typescript
import { AtlasMesh } from '@atlas402/mesh';

const mesh = new AtlasMesh({
  facilitatorUrl: 'https://facilitator.payai.network',
  merchantAddress: '0x...',
});

const service = await mesh.registerService({
  name: 'My API Service',
  endpoint: 'https://api.example.com/data',
  category: 'Data',
  price: '0.10',
  network: 'base',
  scheme: 'x402+eip712',
  merchantAddress: '0x...',
});
```

## License

Apache 2.0
