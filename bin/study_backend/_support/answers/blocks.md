# Block explain drills — answer keys

## 01_rest_api_jwt
- statusCreated: 201
- idempotentUpdate: put
- jwtSegmentCount: 3
- jwtVerifyKey: public
- apiVersioning: url (or header)

## 02_databases_sql
- defaultIndexType: b-tree
- compositeOrderMatters: yes
- nPlusOneFix: eager / join / batch
- poolAvoids: handshake
- explainShows: execution plan

## 03_distributed_resilience
- backoffBase: 2
- breakerPrevents: cascading
- idempotencyKeyLocation: header
- capSacrifice: consistency OR availability
- sagaCoordination: orchestration / choreography / events

## 04_realtime_webrtc
- mediaPath: p2p / peer
- signalingCarries: sdp
- iceGathers: candidates
- stunDiscovers: address / ip
- turnWhen: fails / blocked

## 05_workflows_messaging
- airflowUnit: task
- dagMeaning: directed acyclic graph
- airflowIdempotent: idempotent
- dlqFor: poison / failed
- workQueueConsumers: consumers

## 06_devops_aws
- imageLayers: layers
- ciBeforeDeploy: test
- lambdaBestFor: short / event
- cdnCaches: static
- healthProbe: liveness / readiness

## 07_compliance_security
- xsdValidates: structure / schema
- signatureProves: integrity (and authenticity)
- hashAlgo: sha256
- replayPrevention: nonce / timestamp / id
- auditStores: who and when

## 08_go_systems
- goroutineModel: m:n
- contextFor: cancellation / deadline
- torrentEncoding: bencode
- keepAliveReuses: connection
- gracefulShutdown: listener / drain
