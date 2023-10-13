            {% autoescape off %}[$class: 'hudson.model.TextParameterDefinition',
                name: '{{ name }}',
                description: '{{ description }}',
                defaultValue: '{{ defaultValue }}',
            ]{% endautoescape %}