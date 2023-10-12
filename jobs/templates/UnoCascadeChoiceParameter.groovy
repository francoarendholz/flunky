            {% autoescape off %}[$class: 'CascadeChoiceParameter', 
                choiceType: '{{ choiceType }}', 
                description: '{{ description }}', 
                filterLength: {{ filterLength }}, 
                filterable: {{ filterable }}, 
                name: '{{ name }}', 
                referencedParameters: '{{ referencedParameters }}', 
                script: [
                    $class: '{{ scriptClass }}', 
                    fallbackScript: [
                        classpath: [{{ fallbackClassPath }}], 
                        sandbox: {{ fallbackSandbox }}, 
                        script: 
                            '''{{ fallbackScript }}'''
                    ], 
                    script: [
                        classpath: [{{ classPath }}], 
                        sandbox: {{ scriptSandbox }}, 
                        script: 
                            '''{{ script }}'''
                    ]
                ]
            ]{% endautoescape %}