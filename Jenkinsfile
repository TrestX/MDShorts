node {
    aws_ecr_repro_uri="${env.ECR_HOST}"
	credential_id="${env.CREDS_ID}"
    stage ('Clone') {
    checkout scm
    }
	stage ('Workspace') {
	  echo "aws_ecr_repro_uri-----------------: ${aws_ecr_repro_uri}"
	  echo "credential_id---------------------: ${credential_id}"
	  echo "ECR_HOST--------------------------: ${ECR_HOST}"
	  echo "APP_NAME-----------------------:${APP_NAME}"
	  echo "ENVIRONMENT-----------------------:${ENVIRONMENT}"
	  echo "GO ENV-----------------------:${go_env}"
	  IMAGE_NAME="${ECR_HOST}/${APP_NAME}:latest"
	  echo "IMAGE_NAME------------------------: ${IMAGE_NAME}"
	  
	 
    }
	stage ('Docker Build and Push') {
      dir("${WORKSPACE}") {
      sh '''
      $(aws ecr get-login --no-include-email --region us-east-2)
	   IMAGE_NAME="${ECR_HOST}/${APP_NAME}:latest"
       docker build --no-cache -t ${IMAGE_NAME} -f Dockerfile .
	    
			docker push $IMAGE_NAME
	        docker rmi $IMAGE_NAME
			docker image prune -f
			'''
		}
    }
	
	stage ('Deployment') {
	  dir("${WORKSPACE}/kubernetes/${ENVIRONMENT}") {
	  sh '''
	  IMAGE_NAME="${ECR_HOST}/${APP_NAME}:latest"
	  kubectl rollout restart deployment/${APP_NAME} --kubeconfig /home/jenkins/.kube/config-${ENVIRONMENT}
	  sleep 40
	  kubectl get deploy -owide --kubeconfig /home/jenkins/.kube/config-${ENVIRONMENT}
	  '''
	  }
	}
}
