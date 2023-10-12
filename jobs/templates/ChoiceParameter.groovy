            {% autoescape off %}[
                $class: 'ChoiceParameterDefinition',
                name: '{{ name }}',
                description: '{{ description }}',
                choices: '{{ choices }}',
            ]{% endautoescape %}