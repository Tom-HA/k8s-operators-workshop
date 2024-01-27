#!/usr/bin/env bash

doc_files=($(find ${PWD} -not -path "*/archetypes/*" -and -not -name "_index.md" -and -not -name "index.md" -and -name "*.md"))

for file in "${doc_files[@]}"; do
  file_base_name=${file##*/}
  if [[ ${file_base_name:0:2} =~ ^[0-9]{2,2} ]]; then
    file_number=${file_base_name:0:2}
    weight=${file_number#0}
    weight_line="weight: ${weight}"
  fi
  title=$(grep -E "^# \w.*" ${file})
  title_line="title: \"${title#\# }\""


  echo "Appending front matter to: ${file}"
  printf '%s\n%s\n%s\n%s%s' "---" "${title_line}" "${weight_line}" "---" "$(cat ${file})" > ${file}
  sed -i.bak "s/${title}//" ${file}
  rm ${file}.bak
done
