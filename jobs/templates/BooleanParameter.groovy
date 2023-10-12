            {% autoescape off %}[
                $class: 'BooleanParameterDefinition',
                name: '{{ name }}',
                description: '{{ description }}',
                defaultValue: {{ defaultValue }}
            ]{% endautoescape %}