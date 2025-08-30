{{- define "pinnedImages" }}
- name: RELATED_IMAGE_KUBEDIG_SNITCH
  value: "{{ .repo }}/{{.images.kubedigSnitch.image}}:{{.images.kubedigSnitch.tag}}"
- name: RELATED_IMAGE_KUBEDIG
  value: "{{ .repo }}/{{.images.kubedig.image}}:{{.images.kubedig.tag}}"
- name: RELATED_IMAGE_KUBEDIG_INIT
  value: "{{ .repo }}/{{.images.kubedigInit.image}}:{{.images.kubedigInit.tag}}"
- name: RELATED_IMAGE_KUBEDIG_RELAY_SERVER
  value: "{{ .repo }}/{{.images.kubedigRelay.image}}:{{.images.kubedigRelay.tag}}"
- name: RELATED_IMAGE_KUBEDIG_CONTROLLER
  value: "{{ .repo }}/{{.images.kubedigController.image}}:{{.images.kubedigController.tag}}"
{{- end }}

{{- define "operatorImage" }}
{{- if .Values.imagePinning }}
{{- printf "%s/%s:%s" .Values.oci_meta.repo .Values.oci_meta.images.kubedigOperator.image .Values.oci_meta.images.kubedigOperator.tag }}
{{- else if eq .Values.kubedigOperator.image.tag "" }}
{{- printf "%s:%s" .Values.kubedigOperator.image.repository .Chart.Version }}
{{- else }}
{{- printf "%s:%s" .Values.kubedigOperator.image.repository .Values.kubedigOperator.image.tag }}
{{- end }}
{{- end }}