# Enabling shell auto completion

## Bash on MacOS
The kubebuilder completion script for Bash can be generated with kubebuilder completion bash. Sourcing this script in your shell enables kubebuilder completion.

However, the kubebuilder completion script depends on bash-completion v2 and Bash 4.1+, which you thus have to previously install.

### Upgrade Bash
#### Upgrade to  Bash 4.1+. You can check your Bash’s version by running:

`echo BASH_VERSION`

#### To upgrade, you may use homebrew:

`brew install bash`

#### Verify by reloading shell:

`echo $BASH_VERSION $SHELL`

### Install bash-completion

#### You can use Homebrew to perform the installation:

`brew install bash-completion@2`

#### Once installed, go ahead and add the path `/usr/local/bin/bash` in the  `/etc/shells`.

`sudo echo “/usr/local/bin/bash” > /etc/shells`

#### Make sure to use installed shell by current user.

`chsh -s /usr/local/bin/bash`

#### Go ahead and run below command to generate completion bash script.

`kubebuilder completion bash > ~/.kube/kubebuilder_autocompletion`

#### Add following content in ~/.bash_profile

```
# kubebuilder autocompletion
if [ -f /usr/local/share/bash-completion/bash_completion ]; then
. /usr/local/share/bash-completion/bash_completion
fi
source ~/.kube/kubebuilder_autocompletion
```
### Restart terminal for the changes to be reflected.





