local ibc = import 'ibc.jsonnet';

ibc {
  'iopn_777-1'+: {
    key_name: 'signer3',
    accounts: super.accounts + [{
      name: 'signer3',
      coins: '0basetcro',
      mnemonic: '${SIGNER3_MNEMONIC}',
    }],
    genesis+: {
      app_state+: {
        iopn+: {
          params+: {
            ibc_timeout: 0,
          },
        },
      },
    },
  },
  relayer+: {
    chains: [super.chains[0] {
      fee_granter: 'crc16z0herz998946wr659lr84c8c556da55dc34hh',  //signer1
    }] + super.chains[1:],
  },
}
