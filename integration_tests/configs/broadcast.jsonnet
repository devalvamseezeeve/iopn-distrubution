local config = import 'default.jsonnet';

config {
  'iopn_777-1'+: {
    validators: [validator {
      client_config: {
        'broadcast-mode': 'sync',
      },
    } for validator in super.validators],
  },
}
