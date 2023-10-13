            {% autoescape off %}[$class: 'hudson.model.PasswordParameterDefinition',
                name: '{{ name }}',
                description: '{{ description }}',
                defaultValue: '{{ defaultValue }}',
            ]{% endautoescape %}
