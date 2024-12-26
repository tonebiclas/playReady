# C:\Program Files (x86)\Amlogic\Aml_Burn_Tool\V2\license
# bgroupcert.dat
# playReadyKeyBox.bin
# zgpriv.dat

import ctypes

# NEED 32 BIT PYTHON
cert = ctypes.CDLL(
   'C:/Program Files (x86)/Amlogic/Aml_Burn_Tool/V2/genDevCert.dll'
)

generate = cert['_PR_GeneratePlayReaykeybox@0']

generate()
