            {% autoescape off %}[$class: 'ChoiceParameter', 
                choiceType: '{{ choiceType }}', 
                description: '{{ description }}', 
                filterLength: {{ filterLength }}, 
                filterable: {{ filterable }}, 
                name: '{{ name }}', 
                script: [
                    $class: '{{ scriptClass }}', 
                    fallbackScript: [
                        classpath: [{{ fallbackClasspath }}], 
                        sandbox: {{ fallbackSandbox }}, 
                        script: 
                            '''{{ fallbackScript }}'''
                    ], 
                    script: [
                        classpath: [{{ scriptClasspath }}], 
                        sandbox: {{ scriptSandbox }}, 
                        script: 
                            '''{{ script }}'''
                    ]
                ]
            ]{% endautoescape %}