{% autoescape off %}[$class: 'hudson.model.StringParameterDefinition',
    name: '{{ name }}',
    description: '{{ description }}',
    defaultValue: '{{ defaultValue }}',
]{% endautoescape %}