using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Enemy : MonoBehaviour
{
    // Start is called before the first frame update

    public float speed = 5;
    Vector3 dir;

    // ���� ���� �ּ�
    public GameObject explosionFactory;

    void Start()
    {
        
        int randValue = Random.Range(0, 10);
        if (randValue < 3)
        {
            // �÷��̾� ����
            GameObject target = GameObject.Find("Player");
            dir = target.transform.position - transform.position;
            dir.Normalize();
        }
        else
        {
            // �Ʒ� ����
            dir = Vector3.down;
        }
    }

    // Update is called once per frame
    void Update()
    {

        transform.position += dir * speed * Time.deltaTime;
    }

    // �浹 ���� ��
    private void OnCollisionEnter(Collision other)
    {

        // ���� ����Ʈ�� �߻���Ŵ
        GameObject explosion = Instantiate(explosionFactory);

        // ���� ȿ���� ��ġ��Ŵ
        explosion.transform.position = transform.position;

        Destroy(other.gameObject);
        Destroy(gameObject);

        // ���� ���� ������ ���� ǥ��

        // 1. Scene���� scoreManager ��ü ã�ƿ���
        GameObject smObject = GameObject.Find("ScoreManager");

        // 2. ScoreManager ���� ������Ʈ���� ���´�.
        ScoreManager sm = smObject.GetComponent<ScoreManager>();

        // 3. ScoreManager ȣ��
        sm.SetScore(sm.GetScore() + 1);


    }



}
