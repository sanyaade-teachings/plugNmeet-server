# Changelog

## [1.9.0](https://github.com/sanyaade-teachings/plugNmeet-server/compare/v1.8.0...v1.9.0) (2025-02-15)


### Features

* azure token renew ([ab08ebd](https://github.com/sanyaade-teachings/plugNmeet-server/commit/ab08ebd6fc74e9fdaee754cc74367513a891de0e))
* extra_data option for user ([70d032d](https://github.com/sanyaade-teachings/plugNmeet-server/commit/70d032dedfebdeec161bfbaee8ed7a5c2b43544b))
* option to auto generate user_id ([89f5fb7](https://github.com/sanyaade-teachings/plugNmeet-server/commit/89f5fb7758d589d123b03e6e8ae8c4bcaf0ce8a6))
* option to disable `virtual_bg` & `raise_hand` feature + `preferred_lang` option for user ([63bc633](https://github.com/sanyaade-teachings/plugNmeet-server/commit/63bc6330c032917861636338f9928964116772e5))
* option to encrypt chat messages ([8f877e9](https://github.com/sanyaade-teachings/plugNmeet-server/commit/8f877e98fc8eaacdadfebf4ea29f2c21453a74fd))
* option to keep deleted recordings in backup ([7d1ee5d](https://github.com/sanyaade-teachings/plugNmeet-server/commit/7d1ee5db73d3689d7137233ba0dbe48294eb82a6))
* option to limit for selecting number of translation languages ([56e5f68](https://github.com/sanyaade-teachings/plugNmeet-server/commit/56e5f68af54835d454170238dc9fbc5cb1e02009))
* option to set `max_size_whiteboard_file` ([1ae3df8](https://github.com/sanyaade-teachings/plugNmeet-server/commit/1ae3df857bfecdd6a29a10b429be66b583008d3e))
* option to set log level ([816d450](https://github.com/sanyaade-teachings/plugNmeet-server/commit/816d450364a611a442d0c64af9e957f675b694d7))
* pre-upload file ([d329079](https://github.com/sanyaade-teachings/plugNmeet-server/commit/d3290791a470e228efef5ef270f88c65cf59c876))
* replaced websocket with nats ([#550](https://github.com/sanyaade-teachings/plugNmeet-server/issues/550)) ([393e124](https://github.com/sanyaade-teachings/plugNmeet-server/commit/393e124b081cbed3d8b564c38b34fa9b9d4e2e36))


### Bug Fixes

* `ex_user_id` was changing because of pointer ([73a5d98](https://github.com/sanyaade-teachings/plugNmeet-server/commit/73a5d98b2da8ea2e8b792cc4aabfef5e6e3977ee))
* added option `-trimpath` ([6b355d9](https://github.com/sanyaade-teachings/plugNmeet-server/commit/6b355d9a89c668e400fefbcfea4dcfae54e14bee))
* **bug:** Ingress was not working in the new Nats solution. Ref: https://github.com/mynaparrot/plugNmeet-server/discussions/611 ([2fc55a4](https://github.com/sanyaade-teachings/plugNmeet-server/commit/2fc55a4e5c02787962a4e77804cfb83b38c7d975))
* bump nodejs ([2779eab](https://github.com/sanyaade-teachings/plugNmeet-server/commit/2779eab3538a008d693e9a577b07aa78c5640a0e))
* **bump:** bump version ([de8f4dd](https://github.com/sanyaade-teachings/plugNmeet-server/commit/de8f4dd1f0dd5086684738c7d762360cfccd28d8))
* **ci:** added beta release ([e7497e2](https://github.com/sanyaade-teachings/plugNmeet-server/commit/e7497e2a84f3768007ecb1a787feea0ee87b883c))
* **ci:** always-update ([1134c50](https://github.com/sanyaade-teachings/plugNmeet-server/commit/1134c50f91c84c7fc9eabf5e82f9b50ffdc01472))
* **ci:** bump beta version ([b79a0ea](https://github.com/sanyaade-teachings/plugNmeet-server/commit/b79a0eadd93a0e06070f12dc087c4f18ec49247a))
* **ci:** bump beta version ([24c9082](https://github.com/sanyaade-teachings/plugNmeet-server/commit/24c90821b4e9309ec6666ddcc19caf0f6069168d))
* **ci:** removed rebase again ([db27ad5](https://github.com/sanyaade-teachings/plugNmeet-server/commit/db27ad54d16a4b39c9b5347ee2f06a8164a95266))
* **ci:** To rebase again ([e1681fc](https://github.com/sanyaade-teachings/plugNmeet-server/commit/e1681fc79aede64a5a9cc297d7a4505461767eb8))
* **ci:** To rebase again ([161b04e](https://github.com/sanyaade-teachings/plugNmeet-server/commit/161b04ed775d25a873617c229a0a8bfb23989672))
* **ci:** To rebase again ([4a34670](https://github.com/sanyaade-teachings/plugNmeet-server/commit/4a3467056d97d57245219b0f7cb6d406c3a1fd92))
* clean up ([6e12ff1](https://github.com/sanyaade-teachings/plugNmeet-server/commit/6e12ff16e7e5dea4e2c41c64dee5e1e8cdb13053))
* delay of user disconnect event ([ba30ad4](https://github.com/sanyaade-teachings/plugNmeet-server/commit/ba30ad4579f4fbed5969ab4185e673a6e1887551))
* deps update ([93fff85](https://github.com/sanyaade-teachings/plugNmeet-server/commit/93fff85fe110fc79da5a602c51d2f91c8031e321))
* deps update ([0d58c13](https://github.com/sanyaade-teachings/plugNmeet-server/commit/0d58c1323a618e4fe6c2bdcb2fd853409851e3a5))
* **deps:** dependencies update ([ef83e0d](https://github.com/sanyaade-teachings/plugNmeet-server/commit/ef83e0daaecb76dc884f53c4c5187d65619d3968))
* **deps:** deps update ([3d8d4b9](https://github.com/sanyaade-teachings/plugNmeet-server/commit/3d8d4b933cdb6fd8a992ddc672d22c83795a2be6))
* **deps:** update github.com/antoniodipinto/ikisocket digest to e598eba ([d30a6d1](https://github.com/sanyaade-teachings/plugNmeet-server/commit/d30a6d1ea7d9ea789bed378a454382e020fc5e96))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 225c77f ([d65403f](https://github.com/sanyaade-teachings/plugNmeet-server/commit/d65403f1c67982b3aebfd1164d71ca33c0d52223))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 3a12694 ([c199b0d](https://github.com/sanyaade-teachings/plugNmeet-server/commit/c199b0db5461d6e998a0ab5123edf1ade8d12859))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 45c10af ([e7ce6c5](https://github.com/sanyaade-teachings/plugNmeet-server/commit/e7ce6c5ae6a4ceed1a033c614c2c818f8f4f5f0b))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 45c10af ([19e1e2b](https://github.com/sanyaade-teachings/plugNmeet-server/commit/19e1e2b5a1ac6e76c2c34c18a85969a5ff4ee12d))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 4d9a8f6 ([d9ec4b4](https://github.com/sanyaade-teachings/plugNmeet-server/commit/d9ec4b43c9abebe5ee05a59ffc9e01b417090691))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 5a79e5a ([b009cba](https://github.com/sanyaade-teachings/plugNmeet-server/commit/b009cba839ef641cd5a987b42b986589636c055d))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 5f7afa8 ([0e915fc](https://github.com/sanyaade-teachings/plugNmeet-server/commit/0e915fc686d726bcc7061195787eb5d6ab5a25d4))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 68b434a ([31d7870](https://github.com/sanyaade-teachings/plugNmeet-server/commit/31d78702fed93a27438e1407addddcb55b341ca4))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 73eee1f ([05b8d9f](https://github.com/sanyaade-teachings/plugNmeet-server/commit/05b8d9f50a96fbfb6681abed793b0072e91ce295))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 76539e3 ([954552b](https://github.com/sanyaade-teachings/plugNmeet-server/commit/954552bc97da21a00da9ab117f2c1d8b1dc9fa06))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 77d9532 ([8993d81](https://github.com/sanyaade-teachings/plugNmeet-server/commit/8993d8181a32dc8508277ea5f94fa0b36398e974))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to ae4987a ([3822d07](https://github.com/sanyaade-teachings/plugNmeet-server/commit/3822d07bab0d1d3ddf76b529919f70f48597c84c))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to bdb59ed ([f9c289d](https://github.com/sanyaade-teachings/plugNmeet-server/commit/f9c289d23ccacba76454697dc4676f080cffeae3))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to bfdbf83 ([6ab1ea0](https://github.com/sanyaade-teachings/plugNmeet-server/commit/6ab1ea0b85a7e34ae4a38d446f8b479ce1f393fc))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to e298c13 ([bf238b9](https://github.com/sanyaade-teachings/plugNmeet-server/commit/bf238b943478857e3a9139a0b0eb0553c729b954))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to eed083b ([f5d0681](https://github.com/sanyaade-teachings/plugNmeet-server/commit/f5d0681369e71d7e92c8ff6cf1291ab7241eeaaf))
* **deps:** update module github.com/ansrivas/fiberprometheus/v2 to v2.7.0 ([4ea74d1](https://github.com/sanyaade-teachings/plugNmeet-server/commit/4ea74d15e944b8ae2bf7a9d58bdb09f0d8aa9559))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.4.3 ([1f41268](https://github.com/sanyaade-teachings/plugNmeet-server/commit/1f41268f65eb7674907635ef27714c65dad23bac))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.5.2 ([099c763](https://github.com/sanyaade-teachings/plugNmeet-server/commit/099c763673649db73d5bb7e2111fb08a6b6489d5))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.6.0 ([1317a6c](https://github.com/sanyaade-teachings/plugNmeet-server/commit/1317a6cc5c6594c53a94499c981ec38ff041358e))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.6.1 ([e3e5aee](https://github.com/sanyaade-teachings/plugNmeet-server/commit/e3e5aee56e88cec1d6c24de8310a0f29b8ccef1d))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.6.2 ([c3bd6a5](https://github.com/sanyaade-teachings/plugNmeet-server/commit/c3bd6a53933418e201d1b26dcff24033906b15e1))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.6.3 ([c061992](https://github.com/sanyaade-teachings/plugNmeet-server/commit/c0619925465bed6d7350c3457f2aaba962674355))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.7.2 ([118e56a](https://github.com/sanyaade-teachings/plugNmeet-server/commit/118e56a40302ae98795984ce29f51d641b2f6486))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.7.3 ([0c51ba8](https://github.com/sanyaade-teachings/plugNmeet-server/commit/0c51ba89ad3037b71bf5a3f7ae83272811c0a78e))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.4 ([3143831](https://github.com/sanyaade-teachings/plugNmeet-server/commit/3143831cc811bfcb58b9b02a30581d7aee16f2aa))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.5 ([611534e](https://github.com/sanyaade-teachings/plugNmeet-server/commit/611534e766f943935c9238ea13bb8e6bc6229659))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.6 ([fc6df57](https://github.com/sanyaade-teachings/plugNmeet-server/commit/fc6df57253b8ff035cc87e3d76d651e2391a9419))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.7 ([56f5d3b](https://github.com/sanyaade-teachings/plugNmeet-server/commit/56f5d3becfacba5c535cf45615c2fdeaf898274d))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.8 ([703623d](https://github.com/sanyaade-teachings/plugNmeet-server/commit/703623d55bbcd8a2b9c613dbe07446670e3f2ccb))
* **deps:** update module github.com/go-jose/go-jose/v3 to v4 ([9dff829](https://github.com/sanyaade-teachings/plugNmeet-server/commit/9dff829406c1c8f08e9e990ae65ee03ee50a8e46))
* **deps:** update module github.com/goccy/go-json to v0.10.3 ([ee184ef](https://github.com/sanyaade-teachings/plugNmeet-server/commit/ee184effa44cfaba81f2604054c63515a6a8a1ca))
* **deps:** update module github.com/goccy/go-json to v0.10.5 ([93ef56e](https://github.com/sanyaade-teachings/plugNmeet-server/commit/93ef56eeeb28292a1d0409684bc71a624bd414aa))
* **deps:** update module github.com/gofiber/contrib/socketio to v1.0.1 ([2c4c856](https://github.com/sanyaade-teachings/plugNmeet-server/commit/2c4c85641852231f7ac4d54e4c36b3c0714f6409))
* **deps:** update module github.com/gofiber/contrib/socketio to v1.1.1 ([bc23ce7](https://github.com/sanyaade-teachings/plugNmeet-server/commit/bc23ce7d88bd249aa67ab5991f1d7578d91937a3))
* **deps:** update module github.com/gofiber/contrib/socketio to v1.1.2 ([3587f83](https://github.com/sanyaade-teachings/plugNmeet-server/commit/3587f83d511e5fd24531c6430ecc7d6b06eb28b2))
* **deps:** update module github.com/gofiber/contrib/websocket to v1.3.1 ([c7b0fa5](https://github.com/sanyaade-teachings/plugNmeet-server/commit/c7b0fa5b2e287caf756bf9550fdad0a47297417f))
* **deps:** update module github.com/gofiber/contrib/websocket to v1.3.2 ([d4c8623](https://github.com/sanyaade-teachings/plugNmeet-server/commit/d4c8623ef0d22b434d96420742c53b33371240a4))
* **deps:** update module github.com/gofiber/fiber/v2 to v2.52.0 ([1ac9334](https://github.com/sanyaade-teachings/plugNmeet-server/commit/1ac93341c35113e82b12b5b6cabd0d8d2e874907))
* **deps:** update module github.com/gofiber/fiber/v2 to v2.52.5 ([b5bd2ed](https://github.com/sanyaade-teachings/plugNmeet-server/commit/b5bd2ed84d103128f12f434e8abc1c9954ba5768))
* **deps:** update module github.com/gofiber/template/html/v2 to v2.1.0 ([a2440a9](https://github.com/sanyaade-teachings/plugNmeet-server/commit/a2440a99b5b5d3f712cbe038dd6c44714d7f6cc5))
* **deps:** update module github.com/gofiber/template/html/v2 to v2.1.1 ([1843f82](https://github.com/sanyaade-teachings/plugNmeet-server/commit/1843f82c66bea9fda98868374ca72f0d8f8ad0ba))
* **deps:** update module github.com/gofiber/template/html/v2 to v2.1.2 ([f1ad262](https://github.com/sanyaade-teachings/plugNmeet-server/commit/f1ad262b0fb13a2a861eef8c090b06206755280f))
* **deps:** update module github.com/livekit/protocol to v1.27.1 ([373f1ba](https://github.com/sanyaade-teachings/plugNmeet-server/commit/373f1ba67264263816b1fb0a912f935d44781e5e))
* **deps:** update module github.com/livekit/server-sdk-go to v1.1.8 ([fdea49d](https://github.com/sanyaade-teachings/plugNmeet-server/commit/fdea49d6937b355782981beef45f53963dfa0c90))
* **deps:** update module github.com/livekit/server-sdk-go to v2 ([58d4d81](https://github.com/sanyaade-teachings/plugNmeet-server/commit/58d4d8153232e12c701ea1812e3168e8fd740d37))
* **deps:** update module github.com/livekit/server-sdk-go to v2 ([e1e6c44](https://github.com/sanyaade-teachings/plugNmeet-server/commit/e1e6c44dbc5da9afedfb214ec380e360c05de164))
* **deps:** update module github.com/livekit/server-sdk-go to v2 ([4e6b9be](https://github.com/sanyaade-teachings/plugNmeet-server/commit/4e6b9be5e4d21bfcb4308b3df94c82ab619df2b7))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.1.2 ([b6c5c65](https://github.com/sanyaade-teachings/plugNmeet-server/commit/b6c5c653e1a01c819380dbcc6afe52a286f07490))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.1.3 ([6301045](https://github.com/sanyaade-teachings/plugNmeet-server/commit/63010451d76f1ba879567d620d3f85fb6dd1bcc0))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.2.0 ([bbe3618](https://github.com/sanyaade-teachings/plugNmeet-server/commit/bbe3618170edb48f77140115b99703adc8a01cb4))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.3.0 ([8a9ba58](https://github.com/sanyaade-teachings/plugNmeet-server/commit/8a9ba58ee49519033ab750a872dc6b33f0aae135))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.3.1 ([d5df05d](https://github.com/sanyaade-teachings/plugNmeet-server/commit/d5df05db06f3f9644971bf8d75b90bf956fd6dd6))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.4.1 ([89cac3f](https://github.com/sanyaade-teachings/plugNmeet-server/commit/89cac3f4eed56abf6a60ac3ef51df88b62503738))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.4.2 ([d689646](https://github.com/sanyaade-teachings/plugNmeet-server/commit/d6896460b5184218044ba81fd4dfc7f6b9405235))
* **deps:** update module github.com/mynaparrot/plugnmeet-protocol to v0.0.0-20240510063558-304155e97aa4 ([ef48d4a](https://github.com/sanyaade-teachings/plugNmeet-server/commit/ef48d4a7f41ce1819240b8c6e6c7619e65c1f8d4))
* **deps:** update module github.com/mynaparrot/plugnmeet-protocol to v0.0.0-20240514095941-b542c3f97e6d ([460121f](https://github.com/sanyaade-teachings/plugNmeet-server/commit/460121f9ffc63b36269ba13ea1af9d5b743cc222))
* **deps:** update module github.com/nats-io/nkeys to v0.4.8 ([6b213d9](https://github.com/sanyaade-teachings/plugNmeet-server/commit/6b213d95a7abacff4092e25fd3a9e4f8f2a3377e))
* **deps:** update module github.com/redis/go-redis/v9 to v9.5.1 ([19a23ab](https://github.com/sanyaade-teachings/plugNmeet-server/commit/19a23ab0f89819721deea75c3bec17865ad1940b))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.0 ([8d7dc9a](https://github.com/sanyaade-teachings/plugNmeet-server/commit/8d7dc9ad888af695cdcab2562e881d76b73b0a15))
* **deps:** update module github.com/urfave/cli/v2 to v2.27.1 ([7ea18aa](https://github.com/sanyaade-teachings/plugNmeet-server/commit/7ea18aae678c0e7035cadfda59bc19ce6720c14d))
* **deps:** update module github.com/urfave/cli/v2 to v2.27.2 ([ffcae30](https://github.com/sanyaade-teachings/plugNmeet-server/commit/ffcae30a2b74f2932c9699aecc95f0195e213e12))
* **deps:** update module github.com/urfave/cli/v2 to v2.27.3 ([2a090de](https://github.com/sanyaade-teachings/plugNmeet-server/commit/2a090de3b9e482dff0f1461518f69e227684eda3))
* **deps:** update module github.com/urfave/cli/v2 to v2.27.5 ([664e915](https://github.com/sanyaade-teachings/plugNmeet-server/commit/664e9153def94cddcb75c3b1ddcb6f4c810295a9))
* **deps:** upgrade `nats.go` ([bc3a121](https://github.com/sanyaade-teachings/plugNmeet-server/commit/bc3a121a67920d5c082d6d6f6c4d2f570258f9bb))
* evaluate response from recorder ([7b81c4a](https://github.com/sanyaade-teachings/plugNmeet-server/commit/7b81c4a6bba042d2887fbdd1a872349a46850527))
* few updates ([3944420](https://github.com/sanyaade-teachings/plugNmeet-server/commit/394442074371516ac16ccef00a4ad70aa6d08d72))
* missed to add `Replicas` ([fb6cba0](https://github.com/sanyaade-teachings/plugNmeet-server/commit/fb6cba0c453c61797578662b56c2254263a52f5e))
* missing EmptyTimeout & MaxParticipants in webhook ([879d5b1](https://github.com/sanyaade-teachings/plugNmeet-server/commit/879d5b1591553fc3c43f9f11b5ac3a8ad2d618d2))
* missing sid problem ([691eb2c](https://github.com/sanyaade-teachings/plugNmeet-server/commit/691eb2c0c8ba91b77d5caf5dbb1da1dfa46f93e8))
* multiuser whiteboard sync problem ([534c196](https://github.com/sanyaade-teachings/plugNmeet-server/commit/534c19616d1d0d631c470d009033eb801e73792f))
* notification not sent ([15c89f3](https://github.com/sanyaade-teachings/plugNmeet-server/commit/15c89f3e787803250ea63c5c810cd7bc7c8c5680))
* release-please-action ([3223400](https://github.com/sanyaade-teachings/plugNmeet-server/commit/3223400ec426f66fd8bbf2628877e6801105be37))
* remove duplicate tag ([377903f](https://github.com/sanyaade-teachings/plugNmeet-server/commit/377903f6347033da290c381c12540021259d201c))
* showing invalid poll result ([21ed66d](https://github.com/sanyaade-teachings/plugNmeet-server/commit/21ed66dc367a76f592815718d6a545777ccec85d))
* some clean up ([2472ad3](https://github.com/sanyaade-teachings/plugNmeet-server/commit/2472ad3ccabda7a7de0006a5deec60996675342f))
* time zone problem ([5b655f5](https://github.com/sanyaade-teachings/plugNmeet-server/commit/5b655f514bfa7debe88fdf883aea612e95880cca))
* **update:** few clean ups ([3d632ec](https://github.com/sanyaade-teachings/plugNmeet-server/commit/3d632ecd5f4394d5759b2ab8c89facf3ffeb5142))
* use `StatusServiceUnavailable` ([554e1f5](https://github.com/sanyaade-teachings/plugNmeet-server/commit/554e1f5e8167ccf2bb91b796f5ffc2369f598908))

## [1.8.0](https://github.com/mynaparrot/plugNmeet-server/compare/v1.7.8...v1.8.0) (2025-02-01)


### Features

* option to keep deleted recordings in backup ([7d1ee5d](https://github.com/mynaparrot/plugNmeet-server/commit/7d1ee5db73d3689d7137233ba0dbe48294eb82a6))


### Bug Fixes

* deps update ([93fff85](https://github.com/mynaparrot/plugNmeet-server/commit/93fff85fe110fc79da5a602c51d2f91c8031e321))
* **deps:** update module github.com/goccy/go-json to v0.10.5 ([93ef56e](https://github.com/mynaparrot/plugNmeet-server/commit/93ef56eeeb28292a1d0409684bc71a624bd414aa))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.4.2 ([d689646](https://github.com/mynaparrot/plugNmeet-server/commit/d6896460b5184218044ba81fd4dfc7f6b9405235))
* few updates ([3944420](https://github.com/mynaparrot/plugNmeet-server/commit/394442074371516ac16ccef00a4ad70aa6d08d72))
* remove duplicate tag ([377903f](https://github.com/mynaparrot/plugNmeet-server/commit/377903f6347033da290c381c12540021259d201c))

## [1.7.8](https://github.com/mynaparrot/plugNmeet-server/compare/v1.7.7...v1.7.8) (2025-01-18)


### Bug Fixes

* bump nodejs ([2779eab](https://github.com/mynaparrot/plugNmeet-server/commit/2779eab3538a008d693e9a577b07aa78c5640a0e))
* **ci:** added beta release ([e7497e2](https://github.com/mynaparrot/plugNmeet-server/commit/e7497e2a84f3768007ecb1a787feea0ee87b883c))
* **ci:** always-update ([1134c50](https://github.com/mynaparrot/plugNmeet-server/commit/1134c50f91c84c7fc9eabf5e82f9b50ffdc01472))
* **ci:** bump beta version ([b79a0ea](https://github.com/mynaparrot/plugNmeet-server/commit/b79a0eadd93a0e06070f12dc087c4f18ec49247a))
* **ci:** bump beta version ([24c9082](https://github.com/mynaparrot/plugNmeet-server/commit/24c90821b4e9309ec6666ddcc19caf0f6069168d))
* **ci:** removed rebase again ([db27ad5](https://github.com/mynaparrot/plugNmeet-server/commit/db27ad54d16a4b39c9b5347ee2f06a8164a95266))
* **ci:** To rebase again ([e1681fc](https://github.com/mynaparrot/plugNmeet-server/commit/e1681fc79aede64a5a9cc297d7a4505461767eb8))
* **ci:** To rebase again ([161b04e](https://github.com/mynaparrot/plugNmeet-server/commit/161b04ed775d25a873617c229a0a8bfb23989672))
* **ci:** To rebase again ([4a34670](https://github.com/mynaparrot/plugNmeet-server/commit/4a3467056d97d57245219b0f7cb6d406c3a1fd92))
* deps update ([0d58c13](https://github.com/mynaparrot/plugNmeet-server/commit/0d58c1323a618e4fe6c2bdcb2fd853409851e3a5))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.8 ([703623d](https://github.com/mynaparrot/plugNmeet-server/commit/703623d55bbcd8a2b9c613dbe07446670e3f2ccb))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.4.1 ([89cac3f](https://github.com/mynaparrot/plugNmeet-server/commit/89cac3f4eed56abf6a60ac3ef51df88b62503738))
* evaluate response from recorder ([7b81c4a](https://github.com/mynaparrot/plugNmeet-server/commit/7b81c4a6bba042d2887fbdd1a872349a46850527))
* use `StatusServiceUnavailable` ([554e1f5](https://github.com/mynaparrot/plugNmeet-server/commit/554e1f5e8167ccf2bb91b796f5ffc2369f598908))

## [1.7.7](https://github.com/mynaparrot/plugNmeet-server/compare/v1.7.6...v1.7.7) (2024-12-20)


### Bug Fixes

* **bug:** Ingress was not working in the new Nats solution. Ref: https://github.com/mynaparrot/plugNmeet-server/discussions/611 ([2fc55a4](https://github.com/mynaparrot/plugNmeet-server/commit/2fc55a4e5c02787962a4e77804cfb83b38c7d975))
* **deps:** deps update ([3d8d4b9](https://github.com/mynaparrot/plugNmeet-server/commit/3d8d4b933cdb6fd8a992ddc672d22c83795a2be6))
* **deps:** upgrade `nats.go` ([bc3a121](https://github.com/mynaparrot/plugNmeet-server/commit/bc3a121a67920d5c082d6d6f6c4d2f570258f9bb))
* **update:** few clean ups ([3d632ec](https://github.com/mynaparrot/plugNmeet-server/commit/3d632ecd5f4394d5759b2ab8c89facf3ffeb5142))

## [1.7.6](https://github.com/mynaparrot/plugNmeet-server/compare/v1.7.5...v1.7.6) (2024-12-07)


### Bug Fixes

* **deps:** dependencies update ([ef83e0d](https://github.com/mynaparrot/plugNmeet-server/commit/ef83e0daaecb76dc884f53c4c5187d65619d3968))
* **deps:** update module github.com/nats-io/nkeys to v0.4.8 ([6b213d9](https://github.com/mynaparrot/plugNmeet-server/commit/6b213d95a7abacff4092e25fd3a9e4f8f2a3377e))
* missed to add `Replicas` ([fb6cba0](https://github.com/mynaparrot/plugNmeet-server/commit/fb6cba0c453c61797578662b56c2254263a52f5e))

## [1.7.5](https://github.com/mynaparrot/plugNmeet-server/compare/v1.7.4...v1.7.5) (2024-11-25)


### Bug Fixes

* added option `-trimpath` ([6b355d9](https://github.com/mynaparrot/plugNmeet-server/commit/6b355d9a89c668e400fefbcfea4dcfae54e14bee))
* **bump:** bump version ([de8f4dd](https://github.com/mynaparrot/plugNmeet-server/commit/de8f4dd1f0dd5086684738c7d762360cfccd28d8))
* clean up ([6e12ff1](https://github.com/mynaparrot/plugNmeet-server/commit/6e12ff16e7e5dea4e2c41c64dee5e1e8cdb13053))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 4d9a8f6 ([d9ec4b4](https://github.com/mynaparrot/plugNmeet-server/commit/d9ec4b43c9abebe5ee05a59ffc9e01b417090691))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.7.3 ([0c51ba8](https://github.com/mynaparrot/plugNmeet-server/commit/0c51ba89ad3037b71bf5a3f7ae83272811c0a78e))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.7 ([56f5d3b](https://github.com/mynaparrot/plugNmeet-server/commit/56f5d3becfacba5c535cf45615c2fdeaf898274d))
* **deps:** update module github.com/livekit/protocol to v1.27.1 ([373f1ba](https://github.com/mynaparrot/plugNmeet-server/commit/373f1ba67264263816b1fb0a912f935d44781e5e))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.3.1 ([d5df05d](https://github.com/mynaparrot/plugNmeet-server/commit/d5df05db06f3f9644971bf8d75b90bf956fd6dd6))

## [1.7.4](https://github.com/mynaparrot/plugNmeet-server/compare/v1.7.3...v1.7.4) (2024-11-04)


### Bug Fixes

* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.6 ([fc6df57](https://github.com/mynaparrot/plugNmeet-server/commit/fc6df57253b8ff035cc87e3d76d651e2391a9419))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.0 ([8d7dc9a](https://github.com/mynaparrot/plugNmeet-server/commit/8d7dc9ad888af695cdcab2562e881d76b73b0a15))
* missing EmptyTimeout & MaxParticipants in webhook ([879d5b1](https://github.com/mynaparrot/plugNmeet-server/commit/879d5b1591553fc3c43f9f11b5ac3a8ad2d618d2))
* release-please-action ([3223400](https://github.com/mynaparrot/plugNmeet-server/commit/3223400ec426f66fd8bbf2628877e6801105be37))
