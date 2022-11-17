# pre exec
[[ -f /var/lib/server-admin/bash-preexec.sh ]] && source /var/lib/server-admin/bash-preexec.sh
# self

preexec(){
  export sship=`echo $SSH_CLIENT | awk '{print $1}'`
  rip=${sship}
  if [[ "${rip}x" == "x" ]]; then rip="localhost"; fi
  if [[ -x /var/lib/server-admin/server-admin ]]; then
    /var/lib/server-admin/server-admin -udp.client=true -data="003:0:127.0.0.1:${rip}#1#${USER}#$1"
  else
    echo "WARN: unable to record bash history"
  fi

}
